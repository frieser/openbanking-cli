package cmd

import (
	"fmt"
	nordigen "github.com/frieser/nordigen-go-lib/v2"
	"github.com/frieser/openbanking/console"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var accountsCmd = &cobra.Command{
	Use:   "accounts",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		bankId := viper.GetString("bankId")
		secretId := viper.GetString("secretId")
		secretKey := viper.GetString("secretKey")
		requisitionId := viper.GetString("requisitionId")

		t := console.Task("Listing accounts...")
		cli, err := nordigen.NewClient(secretId, secretKey)

		if err != nil {
			t.Fail(err)
			os.Exit(1)
		}
		var r nordigen.Requisition

		if requisitionId == "" {
			r, err = GetAuthorization(cli, bankId)

			if err != nil {
				t.Fail(err)
				os.Exit(1)
			}
		} else {
			r, err = cli.GetRequisition(requisitionId)

			if err != nil {
				t.Fail(err)
				os.Exit(1)
			}
		}
		accounts, err := Accounts(cli, r)

		if err != nil {
			t.Fail(err)
			os.Exit(1)
		}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID & Name", "IBAN", "Status", "Owner", "Balance"})
		table.SetAutoMergeCellsByColumnIndex([]int{0, 1, 2, 3})
		table.SetRowLine(true)

		for _, account := range accounts {
			for _, balance := range account.Balances.Balances {
				table.Append([]string{
					fmt.Sprintf("%s \n (%s)", account.Name, account.Id),
					account.Iban,
					account.Status,
					account.Owner,
					balance.BalanceType + " = " + balance.BalanceAmount.Amount + balance.BalanceAmount.Currency})
			}
		}
		t.Success()
		table.Render()

		if len(accounts) > 0 {
			console.Section(fmt.Sprintf(`
			To list an account transactions use for example: 
	
			openbanking-cli txns -a %s -s %s -a %s

			Or to export them:

			openbanking-cli txns export -a %s -i %s -k %s -f ofx
	
		`, accounts[0].Id, secretId, secretKey, accounts[0].Id, secretId, secretKey))
		}

	},
}

func init() {
	accountsCmd.PersistentFlags().StringP("bankId", "b", "", "Bank's ID")
	accountsCmd.PersistentFlags().StringP("requisitionId", "r", "", "Bank's Authorization ID")

	err := viper.BindPFlags(accountsCmd.PersistentFlags())

	if err != nil {
		panic(err)
	}
}
