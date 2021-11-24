package cmd

import (
	"fmt"
	nordigen "github.com/frieser/nordigen-go-lib/v2"
	"github.com/frieser/openbanking/console"
	"github.com/frieser/openbanking/internal"
	"github.com/google/uuid"
	"github.com/manifoldco/promptui"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const redirect = ":3000"

var rootCmd = &cobra.Command{
	Use:   "openbank",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		console.Section(`
		
		`)

		interactive()
	},
}

func init() {
	rootCmd.AddCommand(countryCmd)
	rootCmd.AddCommand(banksCmd)
	rootCmd.AddCommand(authCmd)
	rootCmd.AddCommand(accountsCmd)
	rootCmd.AddCommand(txnsCmd)
	rootCmd.PersistentFlags().StringP("secretId", "i", "", "Nordigen secretId")
	rootCmd.PersistentFlags().StringP("secretKey", "k", "", "Nordigen secretKey")

	err := viper.BindPFlags(rootCmd.PersistentFlags())

	if err != nil {
		panic(err)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func interactive() {
	prompt := promptui.Prompt{Label: "Nordigen Secret Id", Mask: '*'}
	nordigenSecretId, err := prompt.Run()

	prompt = promptui.Prompt{Label: "Nordigen Secret Key", Mask: '*'}
	nordigensecretKey, err := prompt.Run()

	if err != nil {
		panic(err)
	}

	c, err := chooseCountry()

	if err != nil {
		panic(err)
	}
	cli, err := nordigen.NewClient(nordigenSecretId, nordigensecretKey)
	b, err := chooseBank(cli, c)

	if err != nil {
		panic(err)
	}
	endUserId := uuid.NewString()

	err = chooseMaxHistoricDays(cli, b, endUserId)

	if err != nil {
		log.Fatal(err)
	}
	a, err := chooseAccount(cli, b)

	if err != nil {
		panic(err)
	}
	err = chooseAccountAction(cli, a)

	if err != nil {
		panic(err)
	}
}

func chooseAccountAction(cli *nordigen.Client, a internal.Account) error {
	prompt := promptui.Select{
		Label: "What do you want to do?",
		Items: []string{
			"Print the list of transactions",
			"Export transactions to a file",
		},
	}
	option, _, err := prompt.Run()

	if err != nil {
		return err
	}
	switch option {
	case 0:
		err = ShowTxnTable(cli, a.Id)

		if err != nil {
			return err
		}
		console.Section(fmt.Sprintf(`
			To export an account transactions use: 
	
			openbanking-cli txns export -a %s -t {NORDIGEN_TOKEN} -f ofx
	
		`, a.Id))
	case 1:
		prompt := promptui.Select{
			Label: "In which format do you want to export?",
			Items: []string{
				"json",
				"csv",
				"ofx",
			},
		}
		_, format, err := prompt.Run()

		if err != nil {
			return err
		}
		promptInput := promptui.Prompt{
			Label: "Where do you want to export the file?",
			Default: fmt.Sprintf("./export-%s.%s",
				time.Now().Format(time.RFC3339),
				format),
			AllowEdit: true,
		}
		output, err := promptInput.Run()

		if err != nil {
			return err
		}
		pretty := false

		if format == "json" {
			promptInput = promptui.Prompt{
				Label:     "Do you want to prettify the json?",
				IsConfirm: true,
			}
			result, err := promptInput.Run()

			if err != nil {
				return err
			}
			if result == "yes" {
				pretty = true
			}
		}
		err = exportTxns(cli, a.Id, output, format, pretty)

		if err != nil {
			return err
		}
	}

	return nil
}

func ShowTxnTable(cli *nordigen.Client, accountId string) error {
	txns, err := cli.GetAccountTransactions(accountId)

	if err != nil {
		return err
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Date", "Payee", "Memo", "Outflow", "Inflow"})
	table.SetAutoMergeCells(false)
	table.SetAutoMergeCellsByColumnIndex([]int{0})
	table.SetRowLine(true)

	for _, txn := range txns.Transactions.Booked {
		amount, err := strconv.ParseFloat(txn.TransactionAmount.Amount, 2)

		if err != nil {
			return err
		}
		outflow := txn.TransactionAmount.Amount
		inflow := ""

		if amount > 0 {
			inflow = txn.TransactionAmount.Amount
			outflow = ""
		}

		payee := txn.DebtorName

		if payee == "" {
			info := strings.Split(txn.RemittanceInformationUnstructured, "//")

			if len(info) == 3 {
				payee = info[2]
			}
		}
		data := []string{
			txn.BookingDate, payee,
			fmt.Sprintf("%.30s", txn.RemittanceInformationUnstructured),
			outflow, inflow,
		}
		table.Append(data)
	}
	table.Render() // Send output

	return nil
}

func chooseAccount(cli *nordigen.Client, b internal.Bank) (internal.Account, error) {
	t := console.Task("Waiting for authorization...")
	r, err := GetAuthorization(cli, b.Id)

	if err != nil {
		t.Fail("Fail to get authorization")
		return internal.Account{}, err
	}
	t.Success()
	t = console.Task("Listing accounts...")
	accounts, err := Accounts(cli, r)

	if err != nil {
		t.Fail("Fail to list accounts")
		return internal.Account{}, err
	}
	t.Success()
	prompt := promptui.Select{
		Label: "Choose an account",
		Templates: &promptui.SelectTemplates{
			Details: `
	------ Account Info ------
	Name: {{ .Name}}
	Iban: {{ .Iban}}
	Owner: {{ .Owner}}
	Currency: {{ .Currency}}
	Status: {{ .Status}}
	Created at: {{ .CreateAt}}
	Last Access: {{ .LastAccessed}}
	Balances:
	{{ range .Balances.Balances }}
		Amount: {{ .BalanceAmount.Amount }} {{ .BalanceAmount.Currency }}
	{{ end }}
				`,
		},
		Items: accounts,
	}
	selectedAccount, _, err := prompt.Run()

	return accounts[selectedAccount], err
}

func Accounts(cli *nordigen.Client, r nordigen.Requisition) ([]internal.Account, error) {
	accounts := make([]internal.Account, 0)

	for _, account := range r.Accounts {
		mtdt, err := cli.GetAccountMetadata(account)

		if err != nil {
			return nil, err
		}
		dtls, err := cli.GetAccountDetails(account)

		if err != nil {
			return nil, err
		}
		blnc, err := cli.GetAccountBalances(account)

		if err != nil {
			return nil, err
		}
		accounts = append(accounts, internal.Account{Name: dtls.Account.Product, Id: mtdt.Id,
			Iban: dtls.Account.Iban, Owner: dtls.Account.OwnerName,
			Currency: dtls.Account.Currency, Status: dtls.Account.Status,
			CreateAt: mtdt.Created, LastAccessed: mtdt.LastAccessed,
			Balances: blnc,
		})
	}

	return accounts, nil
}

func GetAuthorization(cli *nordigen.Client, bankId string) (nordigen.Requisition, error) {
	requisition := nordigen.Requisition{
		Redirect:      "http://localhost" + redirect,
		Reference:     strconv.Itoa(int(time.Now().Unix())),
		InstitutionId: bankId,
	}
	r, err := cli.CreateRequisition(requisition)

	if err != nil {
		return nordigen.Requisition{}, err
	}
	go internal.OpenBrowser(r.Link)

	ch := make(chan bool, 1)

	go internal.CatchRedirect(redirect, ch)

	<-ch

	// issue in the api, see Status requisition definition
	//for r.Status.Short == "CR" {
	for r.Status == "CR" {
		r, err = cli.GetRequisition(r.Id)

		if err != nil {

			return nordigen.Requisition{}, err
		}
		time.Sleep(1 * time.Second)
	}

	return r, nil
}

func chooseMaxHistoricDays(cli *nordigen.Client, b internal.Bank, endUserId string) error {
	promptDays := promptui.Prompt{
		Label:   fmt.Sprintf("Select Max historic days (Default: %d)", b.TransactionTotalDays),
		Default: strconv.Itoa(b.TransactionTotalDays),
		Validate: func(s string) error {
			days, err := strconv.Atoi(s)

			if err != nil {
				panic(err)
			}

			if days > b.TransactionTotalDays {
				return fmt.Errorf("bigger than default")
			}

			return nil
		},
	}
	days, err := promptDays.Run()

	if err != nil {
		return err
	}
	_, err = strconv.Atoi(days)

	if err != nil {
		return err
	}
	// TODO create end user agreement if b.TransactionTotalDays != d

	return nil
}

func chooseBank(cli *nordigen.Client, c internal.Country) (internal.Bank, error) {
	countryBanks, err := cli.ListInstitutions(c.Code)

	if err != nil {
		panic(err)
	}
	banks := make([]internal.Bank, 0)

	for _, bank := range countryBanks {
		days, err := strconv.Atoi(bank.TransactionTotalDays)

		if err != nil {
			days = 0
		}
		banks = append(banks, internal.Bank{Name: bank.Name, Id: bank.Id, TransactionTotalDays: days})
	}
	prompt := promptui.Select{
		Label: "Choose Bank",
		Items: banks,
	}
	selectedBank, _, err := prompt.Run()

	return banks[selectedBank], err
}

func chooseCountry() (internal.Country, error) {
	prompt := promptui.Select{
		Label: "Which is your bank country",
		Items: internal.Countries,
	}
	selectedCountry, _, err := prompt.Run()

	return internal.Countries[selectedCountry], err
}
