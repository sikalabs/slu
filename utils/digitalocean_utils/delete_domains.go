package digitalocean_utils

import (
	"context"
	"fmt"

	"github.com/digitalocean/godo"
)

func PrepareAllDomainsDelete(token string) []godo.Domain {
	client := godo.NewFromToken(token)
	domains, _, _ := client.Domains.List(context.TODO(), &godo.ListOptions{
		PerPage: 200,
	})
	fmt.Println("Domains marked for clean up:")
	for _, el := range domains {
		fmt.Println(el.Name)
	}
	return domains
}

func DoAllDomainsDelete(token string, Keys []godo.Domain) {
	var err error
	client := godo.NewFromToken(token)
	for _, el := range Keys {
		_, err = client.Domains.Delete(context.TODO(), el.Name)
		if err != nil {
			fmt.Println(err)
		}
	}
}
