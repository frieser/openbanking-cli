package cmd

import (
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"fmt"
	nordigen "github.com/frieser/nordigen-go-lib/v2"
	"github.com/frieser/openbanking/console"
	"github.com/frieser/openbanking/ofx"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strconv"
	"strings"
	"time"
)

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		secretId := viper.GetString("secretId")
		secretKey := viper.GetString("secretKey")
		accountId := viper.GetString("account")
		output := viper.GetString("output")
		format := viper.GetString("format")
		pretty := viper.GetBool("pretty")

		if output == "" {
			output = fmt.Sprintf("./export-%s.%s",
				time.Now().Format(time.RFC3339),
				format)
		}

		cli, err := nordigen.NewClient(secretId, secretKey)

		if err != nil {
			os.Exit(1)
		}
		err = exportTxns(cli, accountId, output, format, pretty)

		if err != nil {
			os.Exit(1)
		}
	},
}

func exportTxns(cli *nordigen.Client, accountId, output, format string, pretty bool) error {
	t := console.Task(fmt.Sprintf("Export %s account txns to %s format in %s", accountId, format, output))

	txns, err := cli.GetAccountTransactions(accountId)

	if err != nil {
		t.Fail(err)
		return err
	}
	f, err := os.Create(output)

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)
	if err != nil {
		t.Fail(err)
		return err
	}
	switch format {
	case "json":
		jsonTxns, err := json.Marshal(&txns)

		if err != nil {
			t.Fail(err)
			return err
		}
		if pretty {
			jsonTxns, err = json.MarshalIndent(&txns, "", "    ")

			if err != nil {
				t.Fail(err)
				return err
			}
		}
		_, err = f.WriteString(string(jsonTxns))

		if err != nil {
			t.Fail(err)
			return err
		}
		t.Success()

		return nil
	case "csv":
		w := csv.NewWriter(f)

		err := w.Write([]string{"txnId", "referenceId", "bookDate", "valueDate", "amount", "currency", "creditorIban",
			"ultimateCreditor", "debtorName", "debtorIban", "information", "bankTxnCode"})

		if err != nil {
			t.Fail(err)
			return err
		}
		for _, txn := range txns.Transactions.Booked {
			err := w.Write([]string{txn.TransactionId, txn.EntryReference, txn.BookingDate, txn.ValueDate,
				txn.TransactionAmount.Amount, txn.TransactionAmount.Currency, txn.CreditorAccount.Iban,
				txn.UltimateCreditor, txn.DebtorName, txn.DebtorAccount.Iban, txn.RemittanceInformationUnstructured,
				txn.BankTransactionCode})

			if err != nil {
				t.Fail(err)
				return err
			}
		}
		w.Flush()
		t.Success()

		return nil
	case "ynba", "ofx":
		ofxTxn := ofx.TxnList{
			XMLName:          xml.Name{},
			CurDef:           "",
			PayerBank:        "",
			PayerAccount:     "",
			PayerAccountType: "",
			Transactions:     make([]ofx.Txn, 0),
		}

		for _, txn := range txns.Transactions.Booked {
			amount, err := strconv.ParseFloat(txn.TransactionAmount.Amount, 32)

			if err != nil {
				t.Fail(err)
				return err
			}
			payee := txn.DebtorName

			if payee == "" {
				info := strings.Split(txn.RemittanceInformationUnstructured, "//")

				if len(info) == 3 {
					payee = info[2]
				}
			}
			ofxTxn.Transactions = append(ofxTxn.Transactions, ofx.Txn{
				XMLName:          xml.Name{},
				Type:             "",
				DatePosted:       txn.BookingDate,
				DateUser:         "",
				Amount:           amount,
				ID:               txn.TransactionId,
				Payee:            payee,
				Memo:             txn.RemittanceInformationUnstructured,
				PayeeBank:        "",
				PayeeAccount:     txn.DebtorAccount.Iban,
				PayeeAccountType: "",
			})
		}
		err := ofxTxn.Write(f)

		if err != nil {
			t.Fail(err)
			return err
		}
		t.Success()

		return nil
	default:
		t.Fail(fmt.Sprintf("%s format not supported", format))
		f.Close()
		_ = os.Remove(f.Name())

		return fmt.Errorf("%s format not supported", format)
	}
}

func init() {
	exportCmd.PersistentFlags().StringP("format", "f", "json", "output format")
	exportCmd.PersistentFlags().StringP("output", "o", "", "output file")
	exportCmd.PersistentFlags().BoolP("pretty", "p", true, "json pretty output")

	err := viper.BindPFlags(exportCmd.PersistentFlags())

	if err != nil {
		panic(err)
	}
}
