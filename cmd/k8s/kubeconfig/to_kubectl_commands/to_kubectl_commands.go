package to_kubectl_commands

import (
	"os"
	"path"

	kubeconfig_cmd "github.com/sikalabs/slu/cmd/k8s/kubeconfig"
	"github.com/sikalabs/slu/utils/k8s_utils"

	"github.com/spf13/cobra"
)

var CmdFlagPath string

var Cmd = &cobra.Command{
	Use:     "to-kubectl-commands",
	Short:   "Transform current context from kubeconfig to kubectl commands",
	Args:    cobra.NoArgs,
	Aliases: []string{"tok"},
	Run: func(c *cobra.Command, args []string) {
		k8s_utils.KubeconfigToKubectlCommandsOrDie(CmdFlagPath)
	},
}

func init() {
	defaultPath := ""
	homeDir, homeDirErr := os.UserHomeDir()
	if homeDirErr == nil {
		defaultPath = path.Join(homeDir, ".kube", "config")
	}

	kubeconfig_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&CmdFlagPath,
		"path",
		"p",
		defaultPath,
		"Path to kubeconfig file",
	)
	if homeDirErr != nil {
		Cmd.MarkFlagRequired("path")
	}
}
