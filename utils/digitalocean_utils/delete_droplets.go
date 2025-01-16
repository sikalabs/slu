package digitalocean_utils

import (
	"context"
	"fmt"

	"github.com/digitalocean/godo"
)

func PrepareAllDropletsDelete(token string) []godo.Droplet {
	client := godo.NewFromToken(token)
	Droplets, _, _ := client.Droplets.List(context.TODO(), &godo.ListOptions{
		PerPage: 200,
	})
	fmt.Println("Droplets marked for clean up:")
	for _, v := range Droplets {
		fmt.Println(v.Name)
	}
	return Droplets
}

func DoAllDropletsDelete(token string, DropletsForCleanUp []godo.Droplet) {
	var err error
	client := godo.NewFromToken(token)
	for _, v := range DropletsForCleanUp {
		_, err = client.Droplets.Delete(context.TODO(), v.ID)
		if err != nil {
			fmt.Println(err)
		}
	}
}
