package install_prometheus_operator_crd

import (
	parent_cmd "github.com/sikalabs/slu/cmd/scripts/kubernetes"
	"github.com/sikalabs/slu/utils/k8s_scripts"
	"github.com/spf13/cobra"
)

var FlagDry bool
var FlagVersion string

var Cmd = &cobra.Command{
	Use:     "install-prometheus-operator-crd",
	Short:   "Install Prometheus Operator CRDs",
	Aliases: []string{"ipocrd"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		k8s_scripts.InstallPrometheusOperatorCRD(FlagVersion, FlagDry)
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
		&FlagVersion,
		"version",
		"v",
		"latest",
		"Version of Prometheus Operator",
	)
}
