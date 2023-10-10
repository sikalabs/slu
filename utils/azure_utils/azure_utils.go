package azure_utils

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/profiles/2017-03-09/resources/mgmt/subscriptions"
	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/resources"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/olekukonko/tablewriter"
)

func GetSubscriptionID() (string, error) {
	authorizer, err := auth.NewAuthorizerFromCLI()
	if err != nil {
		return "", err
	}

	subscriptionsClient := subscriptions.NewClient()
	subscriptionsClient.Authorizer = authorizer

	subList, err := subscriptionsClient.List(context.Background())
	if err != nil {
		return "", err
	}

	for _, sub := range subList.Values() {
		if sub.State == subscriptions.Enabled {
			return *sub.SubscriptionID, nil
		}
	}

	return "", fmt.Errorf("no enabled subscriptions found in tenant")
}

func GetAllResourcesFromCurrentSubscription() []resources.GenericResourceExpanded {
	// Use Azure CLI authentication
	authorizer, err := auth.NewAuthorizerFromCLI()
	handleError(err)

	// Get the subscription ID
	subscriptionID, err := GetSubscriptionID()
	handleError(err)

	// Create a new instance of the resources client
	resourcesClient := resources.NewClient(subscriptionID)
	resourcesClient.Authorizer = authorizer

	// List resources
	resourcesList, err := resourcesClient.ListComplete(context.Background(), "", "", nil)
	handleError(err)

	// Print the resources
	res := []resources.GenericResourceExpanded{}

	for resourcesList.NotDone() {
		r := resourcesList.Value()
		res = append(res, r)
		err = resourcesList.NextWithContext(context.Background())
		handleError(err)
	}

	return res
}

func PrintAllResourcesFromCurrentSubscription() {
	resources := GetAllResourcesFromCurrentSubscription()

	table := tablewriter.NewWriter(os.Stdout)
	table.SetBorder(false)
	table.SetHeader([]string{
		"Name",
		"Type",
		"Location",
	})
	for _, r := range resources {
		table.Append([]string{*r.Name, *r.Type, *r.Location})
	}
	table.Render()
}

func handleError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
