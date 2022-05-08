package vault_upload

import (
	"fmt"

	"github.com/sikalabs/slu/config"
	"github.com/sikalabs/slu/utils/vault_utils"
)

func GetUploadSecrets() (string, string, string, string, string, error) {
	conf := config.ReadConfig()
	sec := config.ReadSecrets()

	client, err := vault_utils.GetClient(conf.SluVault.Url)
	if err != nil {
		return "", "", "", "", "", err
	}
	client.SetToken(sec.SluVault.Token)

	secret, err := client.Logical().Read("secret/data/slu/upload")
	if err != nil {
		return "", "", "", "", "", err
	}
	if secret == nil {
		return "", "", "", "", "", fmt.Errorf("secret secret/data/slu/upload not found")
	}
	data, ok := secret.Data["data"].(map[string]interface{})
	if !ok {
		return "", "", "", "", "", fmt.Errorf("wrong data")
	}
	accessKey, err := getString(data, "ACCESS_KEY", true)
	if err != nil {
		return "", "", "", "", "", err
	}
	secretKey, err := getString(data, "SECRET_KEY", true)
	if err != nil {
		return "", "", "", "", "", err
	}
	region, err := getString(data, "REGION", false)
	if err != nil {
		return "", "", "", "", "", err
	}
	endpoint, err := getString(data, "ENDPOINT", false)
	if err != nil {
		return "", "", "", "", "", err
	}
	bucketName, err := getString(data, "BUCKET_NAME", true)
	if err != nil {
		return "", "", "", "", "", err
	}
	return accessKey, secretKey, region, endpoint, bucketName, nil
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
