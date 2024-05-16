package create_kubernetes_api_ingress

import (
	parent_cmd "github.com/sikalabs/slu/cmd/scripts/kubernetes"
	"github.com/sikalabs/slu/utils/k8s_scripts"
	"github.com/spf13/cobra"
)

var FlagDry bool
var FlagDomain string

var Cmd = &cobra.Command{
	Use:     "create-kubernetes-api-ingress",
	Short:   "Create Ingress for Kubernetes API (kubernetes.default.svc)",
	Aliases: []string{"ckai"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		k8s_scripts.ApplyKubernetesApiIngress(FlagDomain, FlagDry)
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
		&FlagDomain,
		"domain",
		"d",
		"",
		"Domain for Kubernetes API Ingress",
	)
	Cmd.MarkFlagRequired("domain")
}
