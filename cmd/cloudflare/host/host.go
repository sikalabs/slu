package host

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
	parent_cmd "github.com/sikalabs/slu/cmd/cloudflare"

	"github.com/spf13/cobra"

	"github.com/cloudflare/cloudflare-go"
)

var FlagAll bool
var FlagNoBr bool

var Cmd = &cobra.Command{
	Use:     "host",
	Short:   "host command behind the cloudflare proxy",
	Aliases: []string{"h"},
	Args:    cobra.ExactArgs(1),
	Run: func(c *cobra.Command, args []string) {
		host := args[0]

		// Construct a new API object using a global API key
		api, err := cloudflare.NewWithAPIToken(os.Getenv("CLOUDFLARE_API_TOKEN"))
		if err != nil {
			log.Fatal(err)
		}

		// Most API calls require a Context
		ctx := context.Background()

		// Fetch the zone ID
		zoneID, err := api.ZoneIDByName(host)
		if err != nil {
			log.Fatal(err)
		}

		records, _, err := api.ListDNSRecords(
			ctx,
			cloudflare.ZoneIdentifier(zoneID),
			cloudflare.ListDNSRecordsParams{},
		)
		if err != nil {
			log.Fatal(err)
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetBorder(false)
		table.SetHeader([]string{
			"Name",
			"Type",
			"Value",
		})
		for _, record := range records {
			if FlagAll || record.Name == host {
				table.Append([]string{br(record.Name, 30), record.Type, br(record.Content, 60)})
			}
		}
		table.Render()
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
	Cmd.PersistentFlags().BoolVarP(
		&FlagAll,
		"all",
		"A",
		false,
		"Show all records",
	)
	Cmd.PersistentFlags().BoolVarP(
		&FlagNoBr,
		"no-br",
		"B",
		false,
		"Don't break the long lines",
	)
}

func br(s string, max int) string {
	if FlagNoBr {
		return s
	}
	return strings.Join(splitBy(s, max), "\n")
}

func splitBy(s string, n int) []string {
	var ss []string
	for i := 1; i < len(s); i++ {
		if i%n == 0 {
			ss = append(ss, s[:i])
			s = s[i:]
			i = 1
		}
	}
	ss = append(ss, s)
	return ss
}
