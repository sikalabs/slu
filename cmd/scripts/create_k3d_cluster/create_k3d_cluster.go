package create_k3d_cluster

import (
	"fmt"

	parent_cmd "github.com/sikalabs/slu/cmd/scripts"
	"github.com/sikalabs/slu/utils/sh_utils"
	"github.com/spf13/cobra"
)

var FlagInstallTools bool
var FlagDry bool

var Cmd = &cobra.Command{
	Use:     "create-k3d-cluster",
	Aliases: []string{"ck3dc"},
	Short:   "Create k3d cluster with nginx-ingress, cert-manager and cluster issuer",
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		if FlagInstallTools {
			sh("slu install-bin k3d", FlagDry)
			sh("slu install-bin kubectl", FlagDry)
			sh("slu install-bin helm", FlagDry)
			sh("slu install-bin k9s", FlagDry)
		}
		sh(`k3d cluster create default \
--k3s-arg --disable=traefik@server:0 \
--servers 1 --agents 1 \
--port 80:80@loadbalancer \
--port 443:443@loadbalancer \
--wait`, FlagDry)
		sh("slu scripts kubernetes install-ingress", FlagDry)
		sh("slu scripts kubernetes install-cert-manager", FlagDry)
		sh("slu scripts kubernetes install-cluster-issuer", FlagDry)
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
	Cmd.PersistentFlags().BoolVarP(
		&FlagInstallTools,
		"install-tools",
		"i",
		false,
		"Install k3d, kubetl, helm, k9s",
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
