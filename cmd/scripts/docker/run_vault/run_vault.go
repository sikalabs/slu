package run_vault

import (
	"fmt"
	"strings"

	parent_cmd "github.com/sikalabs/slu/cmd/scripts/docker"
	"github.com/sikalabs/slu/utils/exec_utils"
	"github.com/spf13/cobra"
)

var FlagToken string
var FlagPort string
var FlagDryRun bool

var Cmd = &cobra.Command{
	Use:   "run-vault",
	Short: "docker run --name vault -d -p 8200:8200 -e VAULT_DEV_ROOT_TOKEN_ID=root hashicorp/vault server -dev",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		dockerArgs := []string{
			"run",
			"--name", "vault",
			"-d",
			"--cap-add=IPC_LOCK",
			"-p", FlagPort + ":8200",
			"-e", "VAULT_DEV_ROOT_TOKEN_ID=" + FlagToken,
			"-e", "VAULT_DEV_LISTEN_ADDRESS=0.0.0.0:8200",
			"hashicorp/vault",
			"server", "-dev",
		}

		if FlagDryRun {
			fmt.Printf("docker %s\n", strings.Join(dockerArgs, " "))
			return
		}

		exec_utils.ExecOut("docker", dockerArgs...)
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagToken,
		"token",
		"t",
		"root",
		"Set Vault dev root token (VAULT_DEV_ROOT_TOKEN_ID)",
	)
	Cmd.Flags().StringVarP(
		&FlagPort,
		"port",
		"p",
		"8200",
		"Set host port to publish Vault on",
	)
	Cmd.Flags().BoolVar(
		&FlagDryRun,
		"dry-run",
		false,
		"Print command instead of running it",
	)
}
