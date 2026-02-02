package unseal_from_1password

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"

	parentcmd "github.com/sikalabs/slu/cmd/vault"
	"github.com/sikalabs/slu/pkg/utils/op_utils"
	"github.com/spf13/cobra"
)

var FlagVaultGroup string
var FlagVaultName string

type vaultInit struct {
	UnsealKeysB64 []string `json:"unseal_keys_b64"`
}

var Cmd = &cobra.Command{
	Use:   "unseal-from-1password",
	Short: "Unseal Vault pods using keys from 1Password",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		op_utils.CheckOpBinaryExistsOrDie()

		nameRegex := regexp.MustCompile(`^[a-zA-Z0-9_.\-]+$`)
		if !nameRegex.MatchString(FlagVaultGroup) {
			fmt.Fprintf(os.Stderr, "Error: --vault-group must contain only letters, numbers, _, - or .\n")
			os.Exit(1)
		}
		if !nameRegex.MatchString(FlagVaultName) {
			fmt.Fprintf(os.Stderr, "Error: --vault-name must contain only letters, numbers, _, - or .\n")
			os.Exit(1)
		}

		// Download vault init JSON from 1Password
		fileName := FlagVaultName + ".json"
		tmpDir, err := os.MkdirTemp("", "slu-vault-unseal-*")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating temp dir: %v\n", err)
			os.Exit(1)
		}
		tmpFile := filepath.Join(tmpDir, fileName)
		op_utils.GetFileFrom1PasswordOrDie(FlagVaultGroup, fileName, tmpFile)

		// Parse the vault init JSON
		data, err := os.ReadFile(tmpFile)
		os.RemoveAll(tmpDir)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
			os.Exit(1)
		}

		var initData vaultInit
		err = json.Unmarshal(data, &initData)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
			os.Exit(1)
		}

		if len(initData.UnsealKeysB64) == 0 {
			fmt.Fprintf(os.Stderr, "Error: no unseal keys found in vault init JSON\n")
			os.Exit(1)
		}

		// Unseal vault-0, vault-1, vault-2
		for i := 0; i < 3; i++ {
			pod := fmt.Sprintf("vault-%d", i)
			fmt.Printf("Unsealing %s ...\n", pod)
			for _, key := range initData.UnsealKeysB64 {
				cmd := exec.Command(
					"kubectl", "exec", pod,
					"-n", "vault",
					"--",
					"vault", "operator", "unseal", key,
				)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err = cmd.Run()
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error unsealing %s: %v\n", pod, err)
					os.Exit(1)
				}
			}
			fmt.Printf("Successfully unsealed %s\n", pod)
		}
	},
}

func init() {
	parentcmd.Cmd.AddCommand(Cmd)
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
		"1Password vault name (document stored as op://<vault-group>/<name>.json/<name>.json)",
	)
	Cmd.MarkFlagRequired("vault-name")
}
