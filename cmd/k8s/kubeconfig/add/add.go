package add

import (
	kubeconfig_cmd "github.com/sikalabs/slu/cmd/k8s/kubeconfig"
	"github.com/sikalabs/slu/utils/k8s"

	"github.com/spf13/cobra"
)

var CmdFlagPath string

var Cmd = &cobra.Command{
	Use:   "add",
	Short: "Add config to ~/.kube/config",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		k8s.AddToKubeconfigShell(CmdFlagPath)
	},
}

func init() {
	kubeconfig_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&CmdFlagPath,
		"path",
		"p",
		"",
		"New kubeconfig file",
	)
	Cmd.MarkFlagRequired("path")
}
