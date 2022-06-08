package install_cluster_issuer

import (
	parent_cmd "github.com/sikalabs/slu/cmd/scripts/kubernetes"
	"github.com/sikalabs/slu/utils/sh_utils"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "install-cluster-issuer",
	Short:   "Install Let's Encrypt Cluster Issuer",
	Aliases: []string{"ici"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		sh(`cat <<EOF | kubectl apply -f -
apiVersion: cert-manager.io/v1
kind: ClusterIssuer
metadata:
  name: letsencrypt
spec:
  acme:
    email: lets-encrypt-slu@sikamail.com
    server: https://acme-v02.api.letsencrypt.org/directory
    privateKeySecretRef:
      name: letsencrypt-issuer-account-key
    solvers:
    - http01:
        ingress:
          class: nginx
EOF`)
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
