package setup

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	parentcmd "github.com/sikalabs/slu/cmd/terraform"
	"github.com/sikalabs/slu/internal/error_utils"
	"github.com/sikalabs/slu/utils/exec_utils"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type CustomToolingConfig struct {
	Get string `json:"Get" yaml:"Get"`
}

type TerraformConfig struct {
	Meta struct {
		SchemaVersion string `json:"SchemaVersion" yaml:"SchemaVersion"`
	} `json:"Meta" yaml:"Meta"`
	GitlabURL              string                         `json:"GitlabURL" yaml:"GitlabURL"`
	ProjectID              string                         `json:"ProjectID" yaml:"ProjectID"`
	StateName              string                         `json:"StateName" yaml:"StateName"`
	VaultAddr              string                         `json:"VaultAddr,omitempty" yaml:"VaultAddr,omitempty"`
	FilesInVault           map[string]string              `json:"FilesInVault,omitempty" yaml:"FilesInVault,omitempty"`
	FilesWithCustomTooling map[string]CustomToolingConfig `json:"FilesWithCustomTooling,omitempty" yaml:"FilesWithCustomTooling,omitempty"`
}

var FlagUsername string
var FlagToken string
var FlagDryRun bool

var Cmd = &cobra.Command{
	Use:   "setup [additional-terraform-args...]",
	Short: "Initialize Terraform with GitLab backend using config file",
	Run: func(c *cobra.Command, args []string) {
		// Read config file - try YAML first, then JSON
		configDir := filepath.Join(".sikalabs", "terraform")
		yamlPath := filepath.Join(configDir, "terraform.yaml")
		jsonPath := filepath.Join(configDir, "terraform.json")

		var configPath string
		var data []byte
		var err error
		var config TerraformConfig

		// Check if YAML config exists
		if _, err = os.Stat(yamlPath); err == nil {
			configPath = yamlPath
			data, err = os.ReadFile(configPath)
			error_utils.HandleError(err, "Failed to read config file")
			err = yaml.Unmarshal(data, &config)
			error_utils.HandleError(err, "Failed to parse YAML config file")
		} else if _, err = os.Stat(jsonPath); err == nil {
			configPath = jsonPath
			data, err = os.ReadFile(configPath)
			error_utils.HandleError(err, "Failed to read config file")
			err = json.Unmarshal(data, &config)
			error_utils.HandleError(err, "Failed to parse JSON config file")
		} else {
			error_utils.HandleError(fmt.Errorf("config file not found"), "Neither terraform.yaml nor terraform.json found in .sikalabs/terraform/")
		}

		// Try to get token from 1Password if not provided via flags
		username := FlagUsername
		token := FlagToken

		if username == "" && token == "" {
			// Extract domain from GitLab URL
			parsedURL, err := url.Parse(config.GitlabURL)
			if err == nil && parsedURL.Host != "" {
				gitlabDomain := parsedURL.Host
				itemName := fmt.Sprintf("GITLAB_TOKEN_TF_STATE_%s", gitlabDomain)

				fmt.Printf("Checking 1Password for token: %s\n", itemName)

				// Try to get token from 1Password
				opToken, err := exec_utils.ExecStr("op", "item", "get", itemName, "--vault", "employee", "--fields", "password", "--reveal")
				if err == nil && opToken != "" {
					// Token found in 1Password - username is "token" for GitLab token auth
					token = strings.TrimSpace(opToken)
					username = "token"
					fmt.Println("Using token from 1Password")
				} else {
					fmt.Println("Token not found in 1Password, will prompt for credentials")
				}
			}
		}

		// Download files from Vault if FilesInVault is defined
		if len(config.FilesInVault) > 0 {
			if config.VaultAddr == "" {
				error_utils.HandleError(fmt.Errorf("vault address is required"), "VaultAddr must be configured in terraform config when FilesInVault is used")
			}

			// Login to Vault
			fmt.Println("Logging in to Vault...")
			if !FlagDryRun {
				err = exec_utils.ExecOut("vault", "login",
					"-address", config.VaultAddr,
					"-method=oidc")
				error_utils.HandleError(err, "Failed to login to Vault")
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

		// Execute custom tooling commands for files if FilesWithCustomTooling is defined
		if len(config.FilesWithCustomTooling) > 0 {
			fmt.Println("Executing custom tooling commands...")
			for fileName, tooling := range config.FilesWithCustomTooling {
				fmt.Printf("  Executing command for %s\n", fileName)
				if tooling.Get == "" {
					error_utils.HandleError(fmt.Errorf("get command is empty"), fmt.Sprintf("Get command is required for file %s", fileName))
				}

				if !FlagDryRun {
					err = exec_utils.ExecShOut(tooling.Get)
					error_utils.HandleError(err, fmt.Sprintf("Failed to execute custom tooling command for file %s", fileName))
				} else {
					fmt.Printf("    Command: %s\n", tooling.Get)
				}
			}
		}

		// Get username and token from stdin or prompt if not already set from 1Password
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
		err = exec_utils.ExecInOut("terraform", cmdArgs...)
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
