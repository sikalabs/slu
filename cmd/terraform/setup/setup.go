package setup

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

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

var FlagUsername string
var FlagToken string
var FlagDryRun bool

var Cmd = &cobra.Command{
	Use:   "setup [additional-terraform-args...]",
	Short: "Initialize Terraform with GitLab backend using config file",
	Run: func(c *cobra.Command, args []string) {
		// Read config file
		configPath := filepath.Join(".sikalabs", "terraform", "terraform.json")
		data, err := os.ReadFile(configPath)
		error_utils.HandleError(err, "Failed to read config file")

		var config TerraformConfig
		err = json.Unmarshal(data, &config)
		error_utils.HandleError(err, "Failed to parse config file")

		// Download files from Vault if FilesInVault is defined
		if len(config.FilesInVault) > 0 {
			if config.VaultAddr == "" {
				error_utils.HandleError(fmt.Errorf("vault address is required"), "VaultAddr must be configured in terraform.json when FilesInVault is used")
			}

			fmt.Println("Downloading files from Vault...")
			for localPath, vaultPath := range config.FilesInVault {
				fmt.Printf("  Downloading %s from %s\n", localPath, vaultPath)

				if !FlagDryRun {
					err = exec_utils.ExecOut(
						"slu", "vault", "copy-file-from-vault",
						"--vault-address", config.VaultAddr,
						"--secret-path", vaultPath,
						"--file-path", localPath,
					)
					error_utils.HandleError(err, fmt.Sprintf("Failed to download file %s from vault", localPath))
				}
			}
		}

		// Get username and token from stdin or prompt if not provided
		username := FlagUsername
		token := FlagToken

		// Check if stdin has data
		fi, err := os.Stdin.Stat()
		isPipe := err == nil && fi.Mode()&os.ModeNamedPipe != 0

		if isPipe {
			// Input is from a pipe - read username and token (one per line)
			scanner := bufio.NewScanner(os.Stdin)

			if username == "" && scanner.Scan() {
				username = strings.TrimSpace(scanner.Text())
			}

			if token == "" && scanner.Scan() {
				token = strings.TrimSpace(scanner.Text())
			}

			if err := scanner.Err(); err != nil {
				error_utils.HandleError(err, "Failed to read from stdin")
			}
		} else {
			// Interactive mode - prompt for missing values
			if username == "" {
				fmt.Print("GitLab Username: ")
				scanner := bufio.NewScanner(os.Stdin)
				if scanner.Scan() {
					username = strings.TrimSpace(scanner.Text())
				}
				if err := scanner.Err(); err != nil {
					error_utils.HandleError(err, "Failed to read username")
				}
			}

			if token == "" {
				fmt.Print("GitLab Token: ")
				scanner := bufio.NewScanner(os.Stdin)
				if scanner.Scan() {
					token = strings.TrimSpace(scanner.Text())
				}
				if err := scanner.Err(); err != nil {
					error_utils.HandleError(err, "Failed to read token")
				}
			}
		}

		// Require username and token
		if username == "" {
			error_utils.HandleError(fmt.Errorf("username is required"), "Username must be provided via --username flag or stdin")
		}

		if token == "" {
			error_utils.HandleError(fmt.Errorf("token is required"), "Token must be provided via --token flag or stdin")
		}

		// Build terraform init command
		backendAddress := fmt.Sprintf("address=%s/api/v4/projects/%s/terraform/state/%s",
			config.GitlabURL, config.ProjectID, config.StateName)
		lockAddress := fmt.Sprintf("lock_address=%s/api/v4/projects/%s/terraform/state/%s/lock",
			config.GitlabURL, config.ProjectID, config.StateName)
		unlockAddress := fmt.Sprintf("unlock_address=%s/api/v4/projects/%s/terraform/state/%s/lock",
			config.GitlabURL, config.ProjectID, config.StateName)
		usernameConfig := fmt.Sprintf("username=%s", username)
		password := fmt.Sprintf("password=%s", token)

		cmdArgs := []string{
			"init",
			fmt.Sprintf("-backend-config=%s", backendAddress),
			fmt.Sprintf("-backend-config=%s", lockAddress),
			fmt.Sprintf("-backend-config=%s", unlockAddress),
			fmt.Sprintf("-backend-config=%s", usernameConfig),
			fmt.Sprintf("-backend-config=%s", password),
			"-backend-config=lock_method=POST",
			"-backend-config=unlock_method=DELETE",
			"-backend-config=retry_wait_min=5",
		}

		// Add additional args
		cmdArgs = append(cmdArgs, args...)

		// Print command for dry-run
		if FlagDryRun {
			fmt.Printf("terraform %s\n", strings.Join(cmdArgs, " "))
			os.Exit(0)
		}

		// Execute terraform init
		err = exec_utils.ExecOut("terraform", cmdArgs...)
		error_utils.HandleError(err, "Failed to execute terraform init")
	},
}

func init() {
	parentcmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagUsername,
		"username",
		"u",
		"",
		"GitLab username (if not provided, will read from stdin)",
	)
	Cmd.Flags().StringVarP(
		&FlagToken,
		"token",
		"t",
		"",
		"GitLab token (if not provided, will read from stdin)",
	)
	Cmd.Flags().BoolVar(
		&FlagDryRun,
		"dry-run",
		false,
		"Print the command without executing it",
	)
}
