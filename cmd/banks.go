package cmd

import (
	"fmt"
	nordigen "github.com/frieser/nordigen-go-lib/v2"
	"github.com/frieser/openbanking/console"
	"github.com/frieser/openbanking/internal"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var banksCmd = &cobra.Command{
	Use:   "banks",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		country := viper.GetString("country")
		secretId := viper.GetString("secretId")
		secretKey := viper.GetString("secretKey")

		t := console.Task(fmt.Sprintf("Listing available %s banks", country))

		for _, c := range internal.Countries {
			if c.Code == country {
				cli, err := nordigen.NewClient(secretId, secretKey)

				if err != nil {
					t.Fail(err)
					os.Exit(1)
				}
				banks, err := cli.ListInstitutions(c.Code)

				if err != nil {
					t.Fail(err)
					os.Exit(1)
				}
				table := tablewriter.NewWriter(os.Stdout)
				table.SetHeader([]string{"Id", "Name", "BIC", "Max. Historic Days", "Countries"})

				for _, b := range banks {
					table.Append([]string{b.Id, b.Name, b.Bic, b.TransactionTotalDays, fmt.Sprintf("%s", b.Countries)})
				}
				t.Success()
				table.Render()

				return
			}
		}
		t.Fail(fmt.Sprintf("No banks found for %s", country))
		os.Exit(1)
	},
}

func init() {
	banksCmd.PersistentFlags().StringP("country", "c", "", "country")

	err := viper.BindPFlags(banksCmd.PersistentFlags())

	if err != nil {
		panic(err)
	}
}
