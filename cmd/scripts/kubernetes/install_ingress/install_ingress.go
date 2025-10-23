package install_ingress

import (
	parent_cmd "github.com/sikalabs/slu/cmd/scripts/kubernetes"
	"github.com/sikalabs/slu/utils/k8s_scripts"
	"github.com/spf13/cobra"
)

var FlagDry bool
var FlagUseProxyProtocol bool
var FlagInstallOnly bool

var Cmd = &cobra.Command{
	Use:     "install-ingress",
	Short:   "Install Ingress Nginx",
	Aliases: []string{"ii"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		k8s_scripts.InstallIngress(FlagUseProxyProtocol, FlagDry, FlagInstallOnly)
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().BoolVar(
		&FlagDry,
		"dry",
		false,
		"Dry run",
	)
	Cmd.Flags().BoolVar(
		&FlagUseProxyProtocol,
		"use-proxy-protocol",
		false,
		"Use Proxy Protocol",
	)
	Cmd.Flags().BoolVar(
		&FlagInstallOnly,
		"install-only",
		false,
		"Use helm install instead of helm upgrade --install",
	)
}
