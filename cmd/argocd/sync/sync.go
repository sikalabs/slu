package sync

import (
	argocd_cmd "github.com/sikalabs/slu/cmd/argocd"
	"github.com/sikalabs/slu/utils/exec_utils"

	"github.com/spf13/cobra"
)

var FlagNamespace string

var Cmd = &cobra.Command{
	Use:   "sync <name>",
	Short: "Sync ArgoCD Application",
	Args:  cobra.ExactArgs(1),
	Run: func(c *cobra.Command, args []string) {
		name := args[0]
		exec_utils.ExecOut(
			"kubectl", "-n", FlagNamespace, "patch", "application", name, "--type", "merge",
			"-p", `{"spec": {"syncPolicy": {"automated": null}}, "operation": {"sync": {}}}`,
		)
	},
}

func init() {
	argocd_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagNamespace,
		"namespace",
		"n",
		"argocd",
		"ArgoCD Namespace",
	)
}
