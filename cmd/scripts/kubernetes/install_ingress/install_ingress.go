package install_ingress

import (
	"fmt"

	parent_cmd "github.com/sikalabs/slu/cmd/scripts/kubernetes"
	"github.com/sikalabs/slu/utils/sh_utils"
	"github.com/spf13/cobra"
)

var FlagDry bool
var FlagUseProxyProtocol bool

var Cmd = &cobra.Command{
	Use:     "install-ingress",
	Short:   "Install Ingress Nginx",
	Aliases: []string{"ii"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		useProxyProtocol := "false"
		if FlagUseProxyProtocol {
			useProxyProtocol = "true"
		}
		sh(`helm upgrade --install \
	ingress-nginx ingress-nginx \
	--repo https://kubernetes.github.io/ingress-nginx \
	--create-namespace \
	--namespace ingress-nginx \
	--set controller.service.type=ClusterIP \
	--set controller.ingressClassResource.default=true \
	--set controller.kind=DaemonSet \
	--set controller.hostPort.enabled=true \
	--set controller.metrics.enabled=true \
	--set controller.config.use-proxy-protocol=`+useProxyProtocol+` \
	--wait`, FlagDry)
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
