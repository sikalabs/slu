package save_vault_init_json_to_1password

import (
	"fmt"
	"os"
	"path/filepath"

	parentcmd "github.com/sikalabs/slu/cmd/vault"
	"github.com/sikalabs/slu/pkg/utils/op_utils"
	"github.com/spf13/cobra"
)

var FlagFile string
var FlagVaultGroup string
var FlagVaultName string

var Cmd = &cobra.Command{
	Use:   "save-vault-init-json-to-1password",
	Short: "Save vault_init.local.json to 1Password",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		vaultGroup := FlagVaultGroup
		vaultName := FlagVaultName

		if vaultGroup == "" || vaultName == "" {
			fmt.Fprintf(os.Stderr, "Error: --vault-group and --vault-name are required\n")
			os.Exit(1)
		}

		// Copy the file to a temporary file with the desired name (name.json)
		fileName := vaultName + ".json"
		tmpDir, err := os.MkdirTemp("", "slu-vault-init-*")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating temp dir: %v\n", err)
			os.Exit(1)
		}
		defer os.RemoveAll(tmpDir)

		tmpFile := filepath.Join(tmpDir, fileName)
		data, err := os.ReadFile(FlagFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
			os.Exit(1)
		}
		err = os.WriteFile(tmpFile, data, 0600)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing temp file: %v\n", err)
			os.Exit(1)
		}

		// Save to 1Password under the specified vault group
		// This stores as op://<vault-group>/<name.json>/<name.json>
		op_utils.SaveFileTo1PasswordOrDie(vaultGroup, tmpFile)
	},
}

func init() {
	parentcmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagFile,
		"file",
		"f",
		"",
		"Path to vault_init.local.json file",
	)
	Cmd.MarkFlagRequired("file")
	Cmd.Flags().StringVar(
		&FlagVaultGroup,
		"vault-group",
		"",
		"1Password vault group",
	)
	Cmd.MarkFlagRequired("vault-group")
	Cmd.Flags().StringVar(
		&FlagVaultName,
		"vault-name",
		"",
		"1Password vault name (stored as op://<vault-group>/<name>.json/<name>.json)",
	)
	Cmd.MarkFlagRequired("vault-name")
}
