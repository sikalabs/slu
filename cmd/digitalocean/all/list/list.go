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
		// co := config.ReadConfig()
		if FlagAllAccounts {
			for _, do := range s.DigitalOcean {
				fmt.Printf("===== Account: %s =====\n", do.Alias)
				digitalocean_utils.ListAll(do.Token)
				fmt.Println("")
			}
		} else {
			// Show single account (from context)
			do := config_utils.GetCurrentDigitalOceanAccount()
			if do == nil {
				log.Fatal("No context found")
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
}
