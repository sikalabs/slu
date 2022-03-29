package vault_utils

import (
	"fmt"

	"github.com/hashicorp/vault/api"
)

func GetClient(url string) (*api.Client, error) {
	var err error
	client, err := api.NewClient(&api.Config{
		Address: url,
	})
	if err != nil {
		return nil, err
	}
	return client, nil
}

func GetTokenFromUserpass(
	client *api.Client,
	user, password string,
) (string, error) {
	var err error
	path := fmt.Sprintf("auth/userpass/login/%s", user)
	secret, err := client.Logical().Write(path, map[string]interface{}{
		"password": password,
	})
	if err != nil {
		return "", err
	}
	token, err := secret.TokenID()
	if err != nil {
		return "", err
	}
	return token, nil
}
