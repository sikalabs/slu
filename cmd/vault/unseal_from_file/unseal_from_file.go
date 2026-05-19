package unseal_from_file

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	parentcmd "github.com/sikalabs/slu/cmd/vault"
	"github.com/spf13/cobra"
)

var FlagFilePath string

type vaultInit struct {
	UnsealKeysB64 []string `json:"unseal_keys_b64"`
}

var Cmd = &cobra.Command{
	Use:   "unseal-from-file",
	Short: "Unseal Vault pods using keys from a local JSON file",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		data, err := os.ReadFile(FlagFilePath)
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
		&FlagFilePath,
		"file",
		"",
		"Path to vault init JSON file",
	)
	Cmd.MarkFlagRequired("file")
}
