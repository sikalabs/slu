package azure_utils

import (
	"context"
	"fmt"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions"
	"github.com/olekukonko/tablewriter"
	"github.com/sikalabs/slu/internal/error_utils"
)

func GetSubscriptionID() (string, error) {
	cred, err := azidentity.NewAzureCLICredential(nil)
	if err != nil {
		return "", err
	}

	client, err := armsubscriptions.NewClient(cred, nil)
	if err != nil {
		return "", err
	}

	pager := client.NewListPager(nil)
	for pager.More() {
		page, err := pager.NextPage(context.Background())
		if err != nil {
			return "", err
		}
		for _, sub := range page.Value {
			if sub.State != nil && *sub.State == armsubscriptions.SubscriptionStateEnabled {
				return *sub.SubscriptionID, nil
			}
		}
	}

	return "", fmt.Errorf("no enabled subscriptions found in tenant")
}

func GetAllResourcesFromCurrentSubscription() []*armresources.GenericResourceExpanded {
	// Use Azure CLI authentication
	cred, err := azidentity.NewAzureCLICredential(nil)
	error_utils.HandleError(err, "Failed to create Azure credential")

	// Get the subscription ID
	subscriptionID, err := GetSubscriptionID()
	error_utils.HandleError(err, "Failed to get subscription ID")

	// Create a new instance of the resources client
	client, err := armresources.NewClient(subscriptionID, cred, nil)
	error_utils.HandleError(err, "Failed to create resources client")

	// List resources
	res := []*armresources.GenericResourceExpanded{}
	pager := client.NewListPager(nil)

	for pager.More() {
		page, err := pager.NextPage(context.Background())
		error_utils.HandleError(err, "Failed to list resources")
		res = append(res, page.Value...)
	}

	return res
}

func PrintAllResourcesFromCurrentSubscription() {
	resources := GetAllResourcesFromCurrentSubscription()

	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{
		"Name",
		"Type",
		"Location",
	})
	for _, r := range resources {
		table.Append([]string{*r.Name, *r.Type, *r.Location})
	}
	table.Render()
}
