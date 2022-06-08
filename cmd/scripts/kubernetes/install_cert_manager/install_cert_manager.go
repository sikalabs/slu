package install_cert_manager

import (
	parent_cmd "github.com/sikalabs/slu/cmd/scripts/kubernetes"
	"github.com/sikalabs/slu/utils/sh_utils"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "install-cert-manager",
	Short:   "Install Cert-Manager",
	Aliases: []string{"icm"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		sh(`helm upgrade --install \
	cert-manager cert-manager \
	--repo https://charts.jetstack.io \
	--create-namespace \
	--namespace cert-manager \
	--set installCRDs=true \
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
