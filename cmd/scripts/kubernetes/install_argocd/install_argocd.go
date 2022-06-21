package install_argocd

import (
	parent_cmd "github.com/sikalabs/slu/cmd/scripts/kubernetes"
	"github.com/sikalabs/slu/utils/sh_utils"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "install-argocd",
	Short:   "Install ArgoCD",
	Aliases: []string{"iacd"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		sh(`helm upgrade --install \
	argocd argo-cd \
	--repo https://argoproj.github.io/argo-helm \
	--create-namespace \
	--namespace argocd \
	--wait`)
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}

func sh(script string) {
	err := sh_utils.ExecShOutDir("", script)
	if err != nil {
		sh_utils.HandleError(err)
	}
}
