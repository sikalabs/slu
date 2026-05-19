package run_loggen

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	parent_cmd "github.com/sikalabs/slu/cmd/scripts/kubernetes"
	"github.com/sikalabs/slu/utils/time_utils"
	"github.com/spf13/cobra"
)

var FlagDryRun bool

var Cmd = &cobra.Command{
	Use:   "run-loggen",
	Short: "Run slu loggen in Kubernetes",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		kubectlRunArgs := []string{
			"loggen-" + strconv.Itoa(time_utils.Unix()),
			"--image", "ghcr.io/sikalabs/slu:v0.100.0",
		}

		kubectlArgs := append([]string{"run"}, kubectlRunArgs...)
		kubectlArgs = append(kubectlArgs, "--", "slu", "loggen", "--json")

		if FlagDryRun {
			fmt.Printf(
				"kubectl %s\n",
				strings.Join(kubectlArgs, " "),
			)
		} else {
			cmd := exec.Command("kubectl", kubectlArgs...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Stdin = os.Stdin
			cmd.Run()
		}
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().BoolVar(
		&FlagDryRun,
		"dry-run",
		false,
		"Print command instead of running it",
	)
}
