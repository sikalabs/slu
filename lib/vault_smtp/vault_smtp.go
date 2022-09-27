package vault_smtp

import (
	"fmt"
	"strconv"

	"github.com/sikalabs/slu/config"
	"github.com/sikalabs/slu/utils/vault_utils"
)

func GetSMTPSecrets(key string) (string, int, string, string, error) {
	conf := config.ReadConfig()
	sec := config.ReadSecrets()

	client, err := vault_utils.GetClient(conf.SluVault.Url)
	if err != nil {
		return "", 0, "", "", err
	}
	client.SetToken(sec.SluVault.Token)

	secret, err := client.Logical().Read("secret/data/slu/smtp/" + key)
	if err != nil {
		return "", 0, "", "", err
	}
	if secret == nil {
		return "", 0, "", "", fmt.Errorf("secret secret/data/slu/smtp/" + key + " not found")
	}
	data, ok := secret.Data["data"].(map[string]interface{})
	if !ok {
		return "", 0, "", "", fmt.Errorf("wrong data")
	}
	host, err := getString(data, "HOST", true)
	if err != nil {
		return "", 0, "", "", err
	}
	port, err := getIntFromString(data, "PORT", true)
	if err != nil {
		return "", 0, "", "", err
	}
	username, err := getString(data, "USERNAME", false)
	if err != nil {
		return "", 0, "", "", err
	}
	password, err := getString(data, "PASSWORD", false)
	if err != nil {
		return "", 0, "", "", err
	}
	return host, port, username, password, nil
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

func getIntFromString(data map[string]interface{}, key string, required bool) (int, error) {
	val, ok := data[key]
	if !ok {
		if !required {
			return 0, nil
		}
		return 0, fmt.Errorf("key \"%s\" not found", key)
	}
	if val == nil {
		if !required {
			return 0, nil
		}
		return 0, fmt.Errorf("no value for key \"%s\"", key)
	}
	return strconv.Atoi(val.(string))
}
