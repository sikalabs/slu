package save_files_to_vault

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	parentcmd "github.com/sikalabs/slu/cmd/terraform"
	"github.com/sikalabs/slu/internal/error_utils"
	"github.com/sikalabs/slu/utils/exec_utils"
	"github.com/spf13/cobra"
)

type TerraformConfig struct {
	Meta struct {
		SchemaVersion string `json:"SchemaVersion"`
	} `json:"Meta"`
	GitlabURL    string            `json:"GitlabURL"`
	ProjectID    string            `json:"ProjectID"`
	StateName    string            `json:"StateName"`
	VaultAddr    string            `json:"VaultAddr,omitempty"`
	FilesInVault map[string]string `json:"FilesInVault,omitempty"`
}

var FlagDryRun bool

var Cmd = &cobra.Command{
	Use:   "save-files-to-vault",
	Short: "Save files to Vault as defined in FilesInVault configuration",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		// Read config file
		configPath := filepath.Join(".sikalabs", "terraform", "terraform.json")
		data, err := os.ReadFile(configPath)
		error_utils.HandleError(err, "Failed to read config file")

		var config TerraformConfig
		err = json.Unmarshal(data, &config)
		error_utils.HandleError(err, "Failed to parse config file")

		// Check if FilesInVault is configured
		if len(config.FilesInVault) == 0 {
			fmt.Println("No files configured in FilesInVault, nothing to save")
			return
		}

		// Check if VaultAddr is configured
		if config.VaultAddr == "" {
			error_utils.HandleError(fmt.Errorf("vault address is required"), "VaultAddr must be configured in terraform.json when FilesInVault is used")
		}

		// Login to Vault
		fmt.Println("Logging in to Vault...")
		if !FlagDryRun {
			err = exec_utils.ExecOut("vault", "login",
				"-address", config.VaultAddr,
				"-method=oidc")
			error_utils.HandleError(err, "Failed to login to Vault")
		}

		// Upload files to Vault
		fmt.Println("Uploading files to Vault...")
		for localPath, vaultPath := range config.FilesInVault {
			fmt.Printf("  Uploading %s to %s\n", localPath, vaultPath)

			if !FlagDryRun {
				err = exec_utils.ExecOut(
					"slu", "vault", "copy-file-to-vault",
					"--vault-address", config.VaultAddr,
					"--secret-path", vaultPath,
					"--file-path", localPath,
				)
				error_utils.HandleError(err, fmt.Sprintf("Failed to upload file %s to vault", localPath))
			}
		}

		fmt.Println("All files saved to Vault successfully")
	},
}

func init() {
	parentcmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().BoolVar(
		&FlagDryRun,
		"dry-run",
		false,
		"Print the operations without executing them",
	)
}
