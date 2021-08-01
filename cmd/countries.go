package cmd

import (
	"github.com/frieser/openbanking/console"
	"github.com/frieser/openbanking/internal"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"os"
)

var countryCmd = &cobra.Command{
	Use:   "countries",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"", "Name", "Code"})

		for _, c := range internal.Countries {
			table.Append([]string{c.Emoji, c.Name, c.Code})
		}
		table.Render()

		console.Section(`
		To list all the banks in a country use:

		openbanking-cli banks -c {COUNTRY_CODE} -t {NORDIGEN_TOKEN}
		`)
	},
}
