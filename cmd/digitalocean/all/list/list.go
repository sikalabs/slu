package list

import (
	parent_cmd "github.com/sikalabs/slu/cmd/digitalocean/all"
	"github.com/sikalabs/slu/utils/digitalocean_utils"

	"github.com/spf13/cobra"
)

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
		token := digitalocean_utils.GetToken()
		digitalocean_utils.ListAll(token)
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}
