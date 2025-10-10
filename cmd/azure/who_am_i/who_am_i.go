package password

import (
	"context"
	"fmt"
	"log"

	parent_cmd "github.com/sikalabs/slu/cmd/azure"
	"github.com/sikalabs/slu/internal/error_utils"

	"github.com/spf13/cobra"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions"
)

var Cmd = &cobra.Command{
	Use:     "who-am-i",
	Short:   "Get Azure Account Info",
	Aliases: []string{"w"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		printSubscriptionInfo()
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}

func printSubscriptionInfo() {
	// Use Azure CLI credential for authentication
	cred, err := azidentity.NewAzureCLICredential(nil)
	error_utils.HandleError(err, "Failed to create Azure credential")

	// Create a client
	client, err := armsubscriptions.NewClient(cred, nil)
	error_utils.HandleError(err, "Failed to create subscriptions client")

	// Get the list of subscriptions
	pager := client.NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		error_utils.HandleError(err, "Failed to list subscriptions")

		for _, sub := range page.Value {
			if sub.State != nil && *sub.State == armsubscriptions.SubscriptionStateEnabled {
				fmt.Printf("Subscription ID:   %s\n", *sub.SubscriptionID)
				fmt.Printf("Subscription Name: %s\n", *sub.DisplayName)
				return
			}
		}
	}

	log.Fatalln("No active subscription found")
}
