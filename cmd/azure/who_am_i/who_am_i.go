package password

import (
	"context"
	"fmt"
	"log"

	parent_cmd "github.com/sikalabs/slu/cmd/azure"
	"github.com/sikalabs/slu/internal/error_utils"

	"github.com/spf13/cobra"

	"github.com/Azure/azure-sdk-for-go/profiles/2017-03-09/resources/mgmt/subscriptions"
	"github.com/Azure/go-autorest/autorest/azure/auth"
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
	// Use environment variables for authentication
	authorizer, err := auth.NewAuthorizerFromCLI()
	error_utils.HandleError(err, "Failed to create Azure authorizer")

	// Create a client
	subscriptionsClient := subscriptions.NewClient()
	subscriptionsClient.Authorizer = authorizer

	// Get the list of subscriptions
	subList, err := subscriptionsClient.List(context.Background())
	error_utils.HandleError(err, "Failed to list subscriptions")

	for _, sub := range subList.Values() {
		if sub.State == subscriptions.Enabled {
			fmt.Printf("Subscription ID:   %s\n", *sub.SubscriptionID)
			fmt.Printf("Subscription Name: %s\n", *sub.DisplayName)
			return
		}
	}

	log.Fatalln("No active subscription found")
}
