package list

import (
	"fmt"
	"log"

	parent_cmd "github.com/sikalabs/slu/cmd/digitalocean/all"
	"github.com/sikalabs/slu/config"
	"github.com/sikalabs/slu/utils/config_utils"
	"github.com/sikalabs/slu/utils/digitalocean_utils"

	"github.com/spf13/cobra"
)

var FlagAllAccounts bool
var FlagAlias string

var Cmd = &cobra.Command{
	Use:   "list",
	Short: "List all resources in DigitalOcean account",
	Long: `List all resources in DigitalOcean account

Currently, all resources means:

- Kubernetes Clusters
- Droplets
- Volumes
- LoadBalancers
- Domains
- Keys (SSH)
`,
	Args: cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		s := config.ReadSecrets()
		if FlagAllAccounts {
			for _, do := range s.DigitalOcean {
				fmt.Printf("===== Account: %s =====\n", do.Alias)
				digitalocean_utils.ListAll(do.Token)
				fmt.Println("")
			}
		} else {
			var do *config.SluSecretsDigitalOcean
			// Show single account

			if FlagAlias == "" {
				// from context
				do = config_utils.GetCurrentDigitalOceanAccount()
			} else {
				// from flag
				do = config_utils.GetDigitalOceanAccountByAlias(FlagAlias)
			}

			if do == nil {
				log.Fatal("No credentials or context found")
			}

			fmt.Printf("===== Account: %s =====\n", do.Alias)
			digitalocean_utils.ListAll(do.Token)
			fmt.Println("")
		}
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
	Cmd.PersistentFlags().BoolVarP(
		&FlagAllAccounts,
		"all",
		"A",
		false,
		"List resources from all accounts",
	)
	Cmd.PersistentFlags().StringVarP(
		&FlagAlias,
		"alias",
		"a",
		"",
		"Use specific account (instead of selected account in context)",
	)
}
