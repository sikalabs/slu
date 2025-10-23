package copy_file_to_vault

import (
	"encoding/base64"
	"fmt"
	"os"
	"os/exec"

	parentcmd "github.com/sikalabs/slu/cmd/vault"
	"github.com/spf13/cobra"
)

var FlagVaultAddress string
var FlagSecretPath string
var FlagFilePath string

var Cmd = &cobra.Command{
	Use:   "copy-file-to-vault",
	Short: "Upload a file to Vault (base64 encoded)",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		err := vaultFileToVault(FlagVaultAddress, FlagSecretPath, FlagFilePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
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
		&FlagSecretPath,
		"secret-path",
		"s",
		"",
		"Vault secret path",
	)
	Cmd.MarkFlagRequired("secret-path")
	Cmd.Flags().StringVarP(
		&FlagFilePath,
		"file-path",
		"f",
		"",
		"Local file path to upload",
	)
	Cmd.MarkFlagRequired("file-path")
}

func vaultFileToVault(vaultAddr, secretPath, filePath string) error {
	// Read the file
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	// Base64 encode the file data
	encodedData := base64.StdEncoding.EncodeToString(fileData)

	// Execute vault kv put command
	cmd := exec.Command("vault", "kv", "put", secretPath, fmt.Sprintf("data=%s", encodedData))
	cmd.Env = append(os.Environ(), fmt.Sprintf("VAULT_ADDR=%s", vaultAddr))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("vault command failed: %w", err)
	}

	fmt.Printf("File uploaded to Vault: %s\n", secretPath)
	return nil
}
