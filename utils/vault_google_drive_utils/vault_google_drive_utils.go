package vault_google_drive_utils

import (
	"fmt"

	"github.com/sikalabs/slu/config"
	"github.com/sikalabs/slu/utils/vault_utils"
)

func GetGoogleDriveUploadSecrets(vaultPath string) (string, string, error) {
	conf := config.ReadConfig()
	sec := config.ReadSecrets()

	client, err := vault_utils.GetClient(conf.SluVault.Url)
	if err != nil {
		return "", "", err
	}
	client.SetToken(sec.SluVault.Token)

	secret, err := client.Logical().Read(vaultPath)
	if err != nil {
		return "", "", err
	}
	if secret == nil {
		return "", "", fmt.Errorf("secret " + vaultPath + " not found")
	}
	data, ok := secret.Data["data"].(map[string]interface{})
	if !ok {
		return "", "", fmt.Errorf("wrong data")
	}
	clientId, err := getString(data, "CLIENT_ID", true)
	if err != nil {
		return "", "", err
	}
	clientSecret, err := getString(data, "CLIENT_SECRET", true)
	if err != nil {
		return "", "", err
	}
	return clientId, clientSecret, nil
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
