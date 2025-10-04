package read_file

import (
	"fmt"

	"github.com/hashicorp/vault/api"
	parentcmd "github.com/sikalabs/slu/cmd/vault"
	"github.com/sikalabs/slu/internal/error_utils"
	"github.com/sikalabs/slu/utils/vault_utils"
	"github.com/spf13/cobra"
)

var FlagVaultAddress string
var FlagVaultToken string
var FlagVaultPath string
var FlagFilePath string

var Cmd = &cobra.Command{
	Use:   "read-file",
	Short: "Read a file from HashiCorp Vault and save it to disk",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		readFileFromVault(FlagVaultAddress, FlagVaultToken, FlagVaultPath, FlagFilePath)
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
	Cmd.Flags().StringVar(
		&FlagVaultPath,
		"vault-path",
		"",
		"Vault path to read the file from",
	)
	Cmd.MarkFlagRequired("vault-path")
	Cmd.Flags().StringVar(
		&FlagFilePath,
		"file-path",
		"",
		"Path where to save the file",
	)
	Cmd.MarkFlagRequired("file-path")
}

func readFileFromVault(vaultAddress, vaultToken, vaultPath, filePath string) {
	client, err := api.NewClient(&api.Config{
		Address: vaultAddress,
	})
	error_utils.HandleError(err, "Failed to create Vault client")
	client.SetToken(vaultToken)
	err = vault_utils.ReadFileFromVault(client, vaultPath, filePath)
	error_utils.HandleError(err, "Failed to read file from vault")
	fmt.Printf("File successfully read from vault and saved to %s\n", filePath)
}
