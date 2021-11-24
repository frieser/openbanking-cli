package cmd

import (
	"fmt"
	nordigen "github.com/frieser/nordigen-go-lib/v2"
	"github.com/frieser/openbanking/console"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		bankId := viper.GetString("bankId")
		secretId := viper.GetString("secretId")
		secretKey := viper.GetString("secretKey")

		cli, err := nordigen.NewClient(secretId, secretKey)

		if err != nil {
			os.Exit(1)
		}
		t := console.Task("Getting Credentials")

		r, err := GetAuthorization(cli, bankId)

		if err != nil {
			t.Fail(err)
			os.Exit(1)
		}
		t.Success()

		console.Section(fmt.Sprintf(`
		Your authorization id : %s

		Now you can run the rest of commands like:

		openbanking-cli accounts -b {YOUR_BANK_ID} -r %s -i %s -k %s
		`, r.Id, r.Id, secretId, secretKey))
	},
}

func init() {
	authCmd.PersistentFlags().StringP("bankId", "b", "", "Bank's ID")

	err := viper.BindPFlags(authCmd.PersistentFlags())

	if err != nil {
		panic(err)
	}
}
