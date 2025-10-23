package install_argocd

import (
	parent_cmd "github.com/sikalabs/slu/cmd/scripts/kubernetes"
	"github.com/sikalabs/slu/utils/k8s_scripts"
	"github.com/spf13/cobra"
)

var FlagDry bool
var FlagNamespace string
var FlagDomain string
var FlagInstallOnly bool

var Cmd = &cobra.Command{
	Use:     "install-argocd",
	Short:   "Install ArgoCD",
	Aliases: []string{"iacd"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		if FlagDomain == "" {
			k8s_scripts.InstallArgoCD(FlagNamespace, FlagDry, FlagInstallOnly)
		} else {
			k8s_scripts.InstallArgoCDDomain(FlagNamespace, FlagDomain, FlagDry, FlagInstallOnly)
		}
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
	Cmd.Flags().StringVarP(
		&FlagDomain,
		"domain",
		"d",
		"",
		"Domain of ArgoCD instance",
	)
	Cmd.Flags().BoolVar(
		&FlagInstallOnly,
		"install-only",
		false,
		"Use helm install instead of helm upgrade --install",
	)
}
