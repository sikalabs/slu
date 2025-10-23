package install_ingress

import (
	parent_cmd "github.com/sikalabs/slu/cmd/scripts/kubernetes"
	"github.com/sikalabs/slu/utils/k8s_scripts"
	"github.com/spf13/cobra"
)

var FlagDry bool
var FlagLoadBalancerIP string
var FlagResourceGroupName string
var FlagInstallOnly bool

var Cmd = &cobra.Command{
	Use:     "install-ingress-aks",
	Short:   "Install Ingress Nginx AKS (Azure)",
	Aliases: []string{"iiaks"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		k8s_scripts.InstallIngressAKS(FlagLoadBalancerIP, FlagResourceGroupName, FlagDry, FlagInstallOnly)
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
		&FlagLoadBalancerIP,
		"loadbalancer-ip",
		"i",
		"",
		"LoadBalancer IP",
	)
	Cmd.MarkFlagRequired("loadbalancer-ip")
	Cmd.Flags().StringVarP(
		&FlagResourceGroupName,
		"resource-group-name",
		"r",
		"",
		"Resource Group Name",
	)
	Cmd.Flags().BoolVar(
		&FlagInstallOnly,
		"install-only",
		false,
		"Use helm install instead of helm upgrade --install",
	)
	Cmd.MarkFlagRequired("resource-group-name")
}
