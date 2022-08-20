package vault_kv

import (
	"fmt"

	"github.com/sikalabs/slu/config"
	"github.com/sikalabs/slu/utils/vault_utils"
)

func Get(path string) (string, error) {
	conf := config.ReadConfig()
	sec := config.ReadSecrets()

	client, err := vault_utils.GetClient(conf.SluVault.Url)
	if err != nil {
		return "", err
	}
	client.SetToken(sec.SluVault.Token)

	secret, err := client.Logical().Read("secret/data/slu/kv/" + path)
	if err != nil {
		return "", err
	}
	if secret == nil {
		return "", fmt.Errorf("secret secret/data/slu/kv/" + path + " not found")
	}
	data, ok := secret.Data["data"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("wrong data")
	}
	value, err := getString(data, "VALUE", true)
	if err != nil {
		return "", err
	}
	return value, nil
}

func Set(path, value string) error {
	conf := config.ReadConfig()
	sec := config.ReadSecrets()

	client, err := vault_utils.GetClient(conf.SluVault.Url)
	if err != nil {
		return err
	}
	client.SetToken(sec.SluVault.Token)

	_, err = client.Logical().Write("secret/data/slu/kv/"+path, map[string]interface{}{
		"data": map[string]string{
			"VALUE": value,
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
