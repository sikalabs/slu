package install_argocd

import (
	"fmt"

	parent_cmd "github.com/sikalabs/slu/cmd/scripts/kubernetes"
	"github.com/sikalabs/slu/utils/sh_utils"
	"github.com/spf13/cobra"
)

var FlagDry bool
var FlagNamespace string

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
	--namespace `+FlagNamespace+` \
	--wait`, FlagDry)
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagNamespace,
		"namespace",
		"n",
		"argocd",
		"Kubernetes Namespace",
	)
	Cmd.Flags().BoolVar(
		&FlagDry,
		"dry",
		false,
		"Dry run",
	)
}

func sh(script string, dry bool) {
	if dry {
		fmt.Println(script)
		return
	}
	err := sh_utils.ExecShOutDir("", script)
	if err != nil {
		sh_utils.HandleError(err)
	}
}
