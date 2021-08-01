package cmd

import (
	"fmt"
	"github.com/frieser/nordigen-go-lib"
	"github.com/frieser/openbanking/console"
	"github.com/google/uuid"
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
		token := viper.GetString("token")

		cli := nordigen.NewClient(token)

		endUserId := uuid.NewString()
		t := console.Task("Getting Credentials")

		r, err := GetAuthorization(cli, bankId, endUserId)

		if err != nil {
			t.Fail(err)
			os.Exit(1)
		}
		t.Success()

		console.Section(fmt.Sprintf(`
		Your authorization id : %s

		Now you can run the rest of commands like:

		openbanking-cli accounts -b {YOUR_BANK_ID} -r %s -t %s
		`, r.Id, r.Id, token))
	},
}

func init() {
	authCmd.PersistentFlags().StringP("bankId", "b", "", "Bank's ID")

	err := viper.BindPFlags(authCmd.PersistentFlags())

	if err != nil {
		panic(err)
	}
}
