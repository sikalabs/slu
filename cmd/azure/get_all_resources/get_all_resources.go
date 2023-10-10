package get_all_resources

import (
	parent_cmd "github.com/sikalabs/slu/cmd/azure"
	"github.com/sikalabs/slu/utils/azure_utils"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "get-all-resources",
	Short:   "Get all resources from Azure Subscription",
	Aliases: []string{"gar"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		azure_utils.PrintAllResourcesFromCurrentSubscription()
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}
