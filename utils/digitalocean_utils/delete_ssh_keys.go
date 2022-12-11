package digitalocean_utils

import (
	"context"
	"fmt"

	"github.com/digitalocean/godo"
)

func PrepareAllSSHKeysDelete(token string) []godo.Key {
	client := godo.NewFromToken(token)
	Keys, _, _ := client.Keys.List(context.TODO(), &godo.ListOptions{})
	fmt.Println("SSH Keys marked for clean up:")
	for _, el := range Keys {
		fmt.Println(el.Name)
	}
	return Keys
}

func DoAllSSHKeysDelete(token string, Keys []godo.Key) {
	var err error
	client := godo.NewFromToken(token)
	for _, el := range Keys {
		_, err = client.Keys.DeleteByID(context.TODO(), el.ID)
		if err != nil {
			fmt.Println(err)
		}
	}
}
