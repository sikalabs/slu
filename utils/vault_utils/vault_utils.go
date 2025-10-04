package vault_utils

import (
	"encoding/base64"
	"fmt"
	"os"

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

func ReadFileFromVault(client *api.Client, vaultPath, filePath string) error {
	secret, err := client.Logical().Read("secret/data/" + vaultPath)
	if err != nil {
		return fmt.Errorf("failed to read from vault: %w", err)
	}
	if secret == nil {
		return fmt.Errorf("secret not found at path: %s", vaultPath)
	}
	data, ok := secret.Data["data"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("invalid data format in vault")
	}
	encodedData, ok := data["data"].(string)
	if !ok {
		return fmt.Errorf("data field not found or invalid in vault")
	}
	decodedData, err := base64.StdEncoding.DecodeString(encodedData)
	if err != nil {
		return fmt.Errorf("failed to decode base64 data: %w", err)
	}
	err = os.WriteFile(filePath, decodedData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}
	return nil
}

func WriteFileToVault(client *api.Client, vaultPath, filePath string) error {
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}
	encodedData := base64.StdEncoding.EncodeToString(fileData)
	_, err = client.Logical().Write("secret/data/"+vaultPath, map[string]interface{}{
		"data": map[string]string{
			"data": encodedData,
		},
	})
	if err != nil {
		return fmt.Errorf("failed to write to vault: %w", err)
	}
	return nil
}
