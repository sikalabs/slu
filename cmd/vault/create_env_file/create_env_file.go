package create_env_file

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/vault/api"
	parentcmd "github.com/sikalabs/slu/cmd/vault"
	"github.com/spf13/cobra"
)

var FlagVaultAddress string
var FlagVaultToken string
var FlagPath string
var FlagMount string
var FlagEnvFileName string

var Cmd = &cobra.Command{
	Use:   "create-env-file",
	Short: "Create .env file form secret stored in HashiCorp Vault",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		createEnvFile(FlagVaultAddress, FlagVaultToken, FlagPath, FlagMount, FlagEnvFileName)
	},
}

func init() {
	parentcmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagVaultAddress,
		"vault-address",
		"a",
		"",
		"Vault address",
	)
	Cmd.MarkFlagRequired("vault-address")
	Cmd.Flags().StringVarP(
		&FlagVaultToken,
		"vault-token",
		"t",
		"",
		"Vault token",
	)
	Cmd.MarkFlagRequired("vault-token")
	Cmd.Flags().StringVarP(
		&FlagPath,
		"path",
		"p",
		"",
		"Path to the secret in Vault",
	)
	Cmd.MarkFlagRequired("path")
	Cmd.Flags().StringVarP(
		&FlagMount,
		"mount",
		"m",
		"secret",
		"Mount point of the secret in Vault",
	)
	Cmd.Flags().StringVarP(
		&FlagEnvFileName,
		"output",
		"o",
		".env",
		"Name of the env file to create",
	)
}

func createEnvFile(vaultAddress, vaultToken, secretPath, mount, envFileName string) {
	// Create a Vault client
	config := api.DefaultConfig()
	config.Address = vaultAddress

	client, err := api.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Vault client: %v", err)
	}

	// Set the token
	client.SetToken(vaultToken)

	// Read the secret from Vault
	secret, err := client.KVv2(mount).Get(context.Background(), secretPath)
	if err != nil {
		log.Fatalf("Failed to read secret from Vault: %v", err)
	}

	// Open the .env file for writing
	file, err := os.Create(envFileName)
	if err != nil {
		log.Fatalf("Failed to create %s file: %v", envFileName, err)
	}
	defer file.Close()

	// Write the environment variables to the .env file
	for key, value := range secret.Data {
		if v, ok := value.(string); ok {
			envLine := fmt.Sprintf("%s=%s\n", key, v)
			_, err := file.WriteString(envLine)
			if err != nil {
				log.Fatalf("Failed to write to %s file: %v", envFileName, err)
			}
		}
	}

	fmt.Printf("%s file created successfully.\n", envFileName)
}
