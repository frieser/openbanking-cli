package cmd

import (
	"fmt"
	nordigen "github.com/frieser/nordigen-go-lib/v2"
	"github.com/frieser/openbanking/console"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var txnsCmd = &cobra.Command{
	Use:   "txns",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		secretId := viper.GetString("secretId")
		secretKey := viper.GetString("secretKey")
		accountId := viper.GetString("account")

		cli, err := nordigen.NewClient(secretId, secretKey)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		t := console.Task(fmt.Sprintf("Listing Transactions for account %s", accountId))
		fmt.Println()
		err = ShowTxnTable(cli, accountId)

		if err != nil {
			t.Fail(err)
			os.Exit(1)
		}
		t.Success()
	},
}

func init() {
	txnsCmd.AddCommand(exportCmd)
	txnsCmd.PersistentFlags().StringP("account", "a", "", "Account ID")

	err := viper.BindPFlags(txnsCmd.PersistentFlags())

	if err != nil {
		panic(err)
	}
}
