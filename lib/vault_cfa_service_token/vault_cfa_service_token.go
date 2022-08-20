package vault_cfa_service_token

import (
	"fmt"

	"github.com/sikalabs/slu/config"
	"github.com/sikalabs/slu/utils/vault_utils"
)

const PATH = "secret/data/slu/cfa_service_token"

func Get(name string) (string, string, error) {
	conf := config.ReadConfig()
	sec := config.ReadSecrets()

	client, err := vault_utils.GetClient(conf.SluVault.Url)
	if err != nil {
		return "", "", err
	}
	client.SetToken(sec.SluVault.Token)

	secret, err := client.Logical().Read(PATH + "/" + name)
	if err != nil {
		return "", "", err
	}
	if secret == nil {
		return "", "", fmt.Errorf("secret " + PATH + "/" + name + " not found")
	}
	data, ok := secret.Data["data"].(map[string]interface{})
	if !ok {
		return "", "", fmt.Errorf("wrong data")
	}
	clientID, err := getString(data, "CLIENT_ID", true)
	if err != nil {
		return "", "", err
	}
	clientSecret, err := getString(data, "CLIENT_SECRET", true)
	if err != nil {
		return "", "", err
	}
	return clientID, clientSecret, nil
}

func Set(name, clientID, clientSecret string) error {
	conf := config.ReadConfig()
	sec := config.ReadSecrets()

	client, err := vault_utils.GetClient(conf.SluVault.Url)
	if err != nil {
		return err
	}
	client.SetToken(sec.SluVault.Token)

	_, err = client.Logical().Write(PATH+"/"+name, map[string]interface{}{
		"data": map[string]string{
			"CLIENT_ID":     clientID,
			"CLIENT_SECRET": clientSecret,
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func getString(data map[string]interface{}, key string, required bool) (string, error) {
	val, ok := data[key]
	if !ok {
		if !required {
			return "", nil
		}
		return "", fmt.Errorf("key \"%s\" not found", key)
	}
	if val == nil {
		if !required {
			return "", nil
		}
		return "", fmt.Errorf("no value for key \"%s\"", key)
	}
	return val.(string), nil
}
