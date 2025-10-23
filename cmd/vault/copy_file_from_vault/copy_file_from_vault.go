package copy_file_from_vault

import (
	"encoding/base64"
	"fmt"
	"os"
	"os/exec"
	"strings"

	parentcmd "github.com/sikalabs/slu/cmd/vault"
	"github.com/spf13/cobra"
)

var FlagVaultAddress string
var FlagSecretPath string
var FlagFilePath string

var Cmd = &cobra.Command{
	Use:   "copy-file-from-vault",
	Short: "Download a file from Vault (base64 decode)",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		err := vaultFileFromVault(FlagVaultAddress, FlagSecretPath, FlagFilePath)
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
		"Local file path to save to",
	)
	Cmd.MarkFlagRequired("file-path")
}

func vaultFileFromVault(vaultAddr, secretPath, filePath string) error {
	// Execute vault kv get command to retrieve the data field
	cmd := exec.Command("vault", "kv", "get", "-field=data", secretPath)
	cmd.Env = append(os.Environ(), fmt.Sprintf("VAULT_ADDR=%s", vaultAddr))
	cmd.Stderr = os.Stderr

	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("vault command failed: %w", err)
	}

	// Base64 decode the data
	encodedData := strings.TrimSpace(string(output))
	decodedData, err := base64.StdEncoding.DecodeString(encodedData)
	if err != nil {
		return fmt.Errorf("failed to decode base64 data: %w", err)
	}

	// Write the decoded data to the file
	err = os.WriteFile(filePath, decodedData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	fmt.Printf("File downloaded from Vault: %s\n", filePath)
	return nil
}
