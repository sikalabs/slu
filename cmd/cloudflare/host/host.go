package host

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	parent_cmd "github.com/sikalabs/slu/cmd/cloudflare"

	"github.com/spf13/cobra"

	"github.com/cloudflare/cloudflare-go"
)

var CmdFlagNamespace string

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
		id, err := getZoneFromDomain(api, host)
		if err != nil {
			log.Fatal(err)
		}

		records, _ := api.DNSRecords(ctx, id, cloudflare.DNSRecord{})
		for _, record := range records {
			if record.Name == host {
				fmt.Printf("%s\t%s\t%s\n", record.Name, record.Type, record.Content)
			}
		}
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}

func getZoneFromDomain(api *cloudflare.API, domain string) (string, error) {
	var err error
	parts := strings.Split(domain, ".")
	for i := len(parts) - 1; i >= 0; i-- {
		id, err := api.ZoneIDByName(strings.Join(parts[i:], "."))
		if err == nil {
			return id, nil
		}
	}
	return "", err
}
