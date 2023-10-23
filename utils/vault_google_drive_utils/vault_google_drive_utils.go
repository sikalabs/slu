package vault_google_drive_utils

import (
	"fmt"
	"log"
	"os"

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

func SetGoogleDriveUploadTokenSecrets(token, vaultPath string) error {
	conf := config.ReadConfig()
	sec := config.ReadSecrets()

	client, err := vault_utils.GetClient(conf.SluVault.Url)
	if err != nil {
		return err
	}
	client.SetToken(sec.SluVault.Token)

	_, err = client.Logical().Write(vaultPath, map[string]interface{}{
		"data": map[string]string{
			"ACCESS_TOKEN": token,
		},
	})
	return err
}

func GetGoogleDriveUploadTokenSecrets(vaultPath string) (string, error) {
	conf := config.ReadConfig()
	sec := config.ReadSecrets()

	client, err := vault_utils.GetClient(conf.SluVault.Url)
	if err != nil {
		return "", err
	}
	client.SetToken(sec.SluVault.Token)

	secret, err := client.Logical().Read(vaultPath)
	if err != nil {
		return "", err
	}
	if secret == nil {
		return "", fmt.Errorf("secret " + vaultPath + " not found")
	}
	data, ok := secret.Data["data"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("wrong data")
	}
	clientId, err := getString(data, "ACCESS_TOKEN", true)
	if err != nil {
		return "", err
	}
	return clientId, nil
}

func GetGoogleDriveUploadTokenSecretsFromVaultOrEnvOrDie() string {
	accessTokenVault,
		_ := GetGoogleDriveUploadTokenSecrets("secret/data/slu/google-drive-upload/token")

	// Client ID
	var accessToken string
	accessTokenEnv := os.Getenv("SLU_GOOGLE_DRIVE_UPLOAD_ACCESS_TOKEN")
	if accessTokenVault != "" {
		accessToken = accessTokenVault
	}
	if accessTokenEnv != "" {
		accessToken = accessTokenEnv
	}
	if accessToken == "" {
		log.Fatalln("SLU_GOOGLE_DRIVE_UPLOAD_ACCESS_TOKEN is empty")
	}
	return accessToken
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
