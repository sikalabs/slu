package ip

import (
	"os"
	"os/exec"
	"strconv"

	"github.com/sikalabs/slu/cmd/root"
	"github.com/sikalabs/slu/utils/time_utils"
	"github.com/spf13/cobra"
)

var FlagImage string
var FlagShell string
var FlagNode string

var Cmd = &cobra.Command{
	Use:   "kdev",
	Short: "Run sikalabs/dev in Kubernetes",
	Args:  cobra.MaximumNArgs(1),
	Run: func(c *cobra.Command, args []string) {
		kubectlRunArgs := []string{
			"dev-" + strconv.Itoa(time_utils.Unix()),
			"--rm", "-ti",
			"--image", FlagImage,
		}

		if FlagNode != "" {
			kubectlRunArgs = append(
				kubectlRunArgs,
				"--overrides", `{"spec": {"nodeName": "`+FlagNode+`"}}`,
			)
		}

		kubectlArgs := append([]string{"run"}, kubectlRunArgs...)
		kubectlArgs = append(kubectlArgs, "--", FlagShell)

		cmd := exec.Command(
			"kubectl", kubectlArgs...,
		)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		cmd.Run()
	},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagImage,
		"image",
		"i",
		"ghcr.io/sikalabs/dev",
		"Container Image",
	)
	Cmd.Flags().StringVarP(
		&FlagShell,
		"shell",
		"s",
		"zsh",
		"Shell to run in container",
	)
	Cmd.Flags().StringVar(
		&FlagNode,
		"node",
		"",
		"Node to run on",
	)
}
