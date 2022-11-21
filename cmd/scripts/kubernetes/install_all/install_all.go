package install_all

import (
	parent_cmd "github.com/sikalabs/slu/cmd/scripts/kubernetes"
	"github.com/sikalabs/slu/utils/k8s_scripts"
	"github.com/spf13/cobra"
)

const DEFAULT_LETS_ENCRYPT_EMAIL = "lets-encrypt-slu@sikamail.com"

var FlagDry bool
var FlagBaseDomain string
var FlagDontUseProxyProtocol bool
var FlagLetsEncryptEmail string
var FlagNoArgoCD bool

var Cmd = &cobra.Command{
	Use:     "install-all",
	Short:   "Install All (Ingress, Cert-Manager, ArgoCD)",
	Aliases: []string{"iall"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		k8s_scripts.InstallIngress(!FlagDontUseProxyProtocol, FlagDry)
		k8s_scripts.InstallCertManager(FlagDry)
		k8s_scripts.InstallClusterIssuer(FlagLetsEncryptEmail, FlagDry)
		if !FlagNoArgoCD {
			k8s_scripts.InstallArgoCDDomain("argocd", "argocd."+FlagBaseDomain, FlagDry)
		}
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
	Cmd.Flags().StringVarP(
		&FlagBaseDomain,
		"base-domain",
		"d",
		"",
		"Base domain of Kubernetes cluster",
	)
	Cmd.MarkFlagRequired("base-domain")
	Cmd.Flags().BoolVar(
		&FlagDontUseProxyProtocol,
		"dont-use-proxy-protocol",
		false,
		"Use Proxy Protocol",
	)
	Cmd.Flags().StringVarP(
		&FlagLetsEncryptEmail,
		"letsencrypt-email",
		"e",
		DEFAULT_LETS_ENCRYPT_EMAIL,
		"Email for Let's Encrypt account & notifications",
	)
	Cmd.Flags().BoolVar(
		&FlagNoArgoCD,
		"no-argocd",
		false,
		"Don't install ArgoCD",
	)
}
