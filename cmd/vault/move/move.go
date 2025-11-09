package move

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	parentcmd "github.com/sikalabs/slu/cmd/vault"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type VaultConfig struct {
	Meta struct {
		SchemaVersion int `yaml:"SchemaVersion"`
	} `yaml:"Meta"`
	VaultAddr string `yaml:"VaultAddr"`
}

var FlagFile string
var FlagDestroy bool

func init() {
	parentcmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(&FlagFile, "file", "f", ".sikalabs/vault/vault.yaml", "Path to vault.yaml configuration file")
	Cmd.Flags().BoolVarP(&FlagDestroy, "destroy", "d", false, "Destroy the old secret (permanent) instead of delete (soft delete)")
}

var Cmd = &cobra.Command{
	Use:     "move <old-path> <new-path>",
	Aliases: []string{"mv"},
	Short:   "Move/rename a secret in Vault",
	Long: `Move/rename a secret in Vault using vault address from .sikalabs/vault/vault.yaml

This command reads the Vault address from the configuration file and moves
a secret from the old path to the new path by:
1. Reading the secret from the old path (vault kv get)
2. Writing it to the new path (vault kv put)
3. Deleting it from the old path (vault kv delete) or destroying it (vault kv destroy) with --destroy flag

The difference between delete and destroy:
- delete: Soft delete, can be undeleted
- destroy: Permanent deletion (destroys all versions and deletes metadata, completely removes from UI)

Example:
  slu vault move secret/old/path secret/new/path
  slu vault mv secret/old/path secret/new/path
  slu vault move -f custom/vault.yaml secret/old/path secret/new/path
  slu vault move --destroy secret/old/path secret/new/path
`,
	Args: cobra.ExactArgs(2),
	Run: func(c *cobra.Command, args []string) {
		oldPath := args[0]
		newPath := args[1]

		err := vaultMove(FlagFile, oldPath, newPath, FlagDestroy)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	},
}

func vaultMove(configFile, oldPath, newPath string, destroy bool) error {
	// Read the YAML file
	data, err := os.ReadFile(configFile)
	if err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	// Parse the YAML
	var config VaultConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return fmt.Errorf("failed to parse YAML: %w", err)
	}

	// Validate VaultAddr
	if config.VaultAddr == "" {
		return fmt.Errorf("VaultAddr not specified in config")
	}

	fmt.Printf("Using Vault at: %s\n", config.VaultAddr)
	fmt.Printf("Moving secret from: %s\n", oldPath)
	fmt.Printf("                to: %s\n", newPath)

	// Step 1: Get the secret from the old path
	fmt.Println("\n1. Reading secret from old path...")
	secretData, err := getVaultSecret(config.VaultAddr, oldPath)
	if err != nil {
		return fmt.Errorf("failed to read secret from old path: %w", err)
	}
	fmt.Printf("   ✓ Secret read successfully (%d key(s))\n", len(secretData))

	// Step 2: Create the secret at the new path
	fmt.Println("\n2. Writing secret to new path...")
	err = putVaultSecret(config.VaultAddr, newPath, secretData)
	if err != nil {
		return fmt.Errorf("failed to write secret to new path: %w", err)
	}
	fmt.Printf("   ✓ Secret written successfully\n")

	// Step 3: Delete or destroy the secret from the old path
	if destroy {
		fmt.Println("\n3. Destroying secret from old path...")
		err = destroyVaultSecret(config.VaultAddr, oldPath)
		if err != nil {
			return fmt.Errorf("failed to destroy secret from old path: %w", err)
		}
		fmt.Printf("   ✓ Secret destroyed successfully (permanent)\n")
	} else {
		fmt.Println("\n3. Deleting secret from old path...")
		err = deleteVaultSecret(config.VaultAddr, oldPath)
		if err != nil {
			return fmt.Errorf("failed to delete secret from old path: %w", err)
		}
		fmt.Printf("   ✓ Secret deleted successfully\n")
	}

	fmt.Printf("\n✓ Secret moved successfully from %s to %s\n", oldPath, newPath)
	return nil
}

func getVaultSecret(vaultAddr, path string) (map[string]interface{}, error) {
	// Execute vault kv get command with JSON output
	cmd := exec.Command("vault", "kv", "get", "-format=json", path)
	cmd.Env = append(os.Environ(), fmt.Sprintf("VAULT_ADDR=%s", vaultAddr))

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("vault get command failed: %w\nStderr: %s", err, stderr.String())
	}

	// Parse the JSON output
	var result struct {
		Data struct {
			Data map[string]interface{} `json:"data"`
		} `json:"data"`
	}

	err = json.Unmarshal(stdout.Bytes(), &result)
	if err != nil {
		return nil, fmt.Errorf("failed to parse vault output: %w", err)
	}

	if len(result.Data.Data) == 0 {
		return nil, fmt.Errorf("no data found in secret")
	}

	return result.Data.Data, nil
}

func putVaultSecret(vaultAddr, path string, data map[string]interface{}) error {
	// Build the vault kv put command
	args := []string{"kv", "put", path}
	for key, value := range data {
		args = append(args, fmt.Sprintf("%s=%v", key, value))
	}

	// Execute the vault command
	cmd := exec.Command("vault", args...)
	cmd.Env = append(os.Environ(), fmt.Sprintf("VAULT_ADDR=%s", vaultAddr))

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("vault put command failed: %w\nStderr: %s", err, stderr.String())
	}

	return nil
}

func deleteVaultSecret(vaultAddr, path string) error {
	// Execute vault kv delete command
	cmd := exec.Command("vault", "kv", "delete", path)
	cmd.Env = append(os.Environ(), fmt.Sprintf("VAULT_ADDR=%s", vaultAddr))

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("vault delete command failed: %w\nStderr: %s", err, stderr.String())
	}

	return nil
}

func destroyVaultSecret(vaultAddr, path string) error {
	// First, get metadata to find all versions
	metaCmd := exec.Command("vault", "kv", "metadata", "get", "-format=json", path)
	metaCmd.Env = append(os.Environ(), fmt.Sprintf("VAULT_ADDR=%s", vaultAddr))

	var metaStdout bytes.Buffer
	var metaStderr bytes.Buffer
	metaCmd.Stdout = &metaStdout
	metaCmd.Stderr = &metaStderr

	err := metaCmd.Run()
	if err != nil {
		return fmt.Errorf("failed to get secret metadata: %w\nStderr: %s", err, metaStderr.String())
	}

	// Parse metadata to get versions
	var metadata struct {
		Data struct {
			Versions map[string]interface{} `json:"versions"`
		} `json:"data"`
	}

	err = json.Unmarshal(metaStdout.Bytes(), &metadata)
	if err != nil {
		return fmt.Errorf("failed to parse metadata: %w", err)
	}

	// Collect all version numbers
	var versions []string
	for version := range metadata.Data.Versions {
		versions = append(versions, version)
	}

	if len(versions) == 0 {
		return fmt.Errorf("no versions found to destroy")
	}

	// Step 1: Destroy all versions
	args := []string{"kv", "destroy", "-versions=" + joinVersions(versions), path}
	cmd := exec.Command("vault", args...)
	cmd.Env = append(os.Environ(), fmt.Sprintf("VAULT_ADDR=%s", vaultAddr))

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("vault destroy command failed: %w\nStderr: %s", err, stderr.String())
	}

	// Step 2: Delete metadata to completely remove from UI
	metaDeleteCmd := exec.Command("vault", "kv", "metadata", "delete", path)
	metaDeleteCmd.Env = append(os.Environ(), fmt.Sprintf("VAULT_ADDR=%s", vaultAddr))

	var metaDeleteStderr bytes.Buffer
	metaDeleteCmd.Stderr = &metaDeleteStderr

	err = metaDeleteCmd.Run()
	if err != nil {
		return fmt.Errorf("vault metadata delete command failed: %w\nStderr: %s", err, metaDeleteStderr.String())
	}

	return nil
}

func joinVersions(versions []string) string {
	result := ""
	for i, v := range versions {
		if i > 0 {
			result += ","
		}
		result += v
	}
	return result
}
