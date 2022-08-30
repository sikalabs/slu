package install_prometheus_operator_crd

import (
	"fmt"

	parent_cmd "github.com/sikalabs/slu/cmd/scripts/kubernetes"
	"github.com/sikalabs/slu/utils/github_utils"
	"github.com/sikalabs/slu/utils/sh_utils"
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
		version := FlagVersion
		if version == "latest" {
			version = github_utils.GetLatestRelease("prometheus-operator", "prometheus-operator")
		}
		sh(`kubectl apply --server-side -f https://raw.githubusercontent.com/prometheus-operator/prometheus-operator/`+version+`/example/prometheus-operator-crd/monitoring.coreos.com_alertmanagerconfigs.yaml`, FlagDry)
		sh(`kubectl apply --server-side -f https://raw.githubusercontent.com/prometheus-operator/prometheus-operator/`+version+`/example/prometheus-operator-crd/monitoring.coreos.com_alertmanagers.yaml`, FlagDry)
		sh(`kubectl apply --server-side -f https://raw.githubusercontent.com/prometheus-operator/prometheus-operator/`+version+`/example/prometheus-operator-crd/monitoring.coreos.com_podmonitors.yaml`, FlagDry)
		sh(`kubectl apply --server-side -f https://raw.githubusercontent.com/prometheus-operator/prometheus-operator/`+version+`/example/prometheus-operator-crd/monitoring.coreos.com_probes.yaml`, FlagDry)
		sh(`kubectl apply --server-side -f https://raw.githubusercontent.com/prometheus-operator/prometheus-operator/`+version+`/example/prometheus-operator-crd/monitoring.coreos.com_prometheuses.yaml`, FlagDry)
		sh(`kubectl apply --server-side -f https://raw.githubusercontent.com/prometheus-operator/prometheus-operator/`+version+`/example/prometheus-operator-crd/monitoring.coreos.com_prometheusrules.yaml`, FlagDry)
		sh(`kubectl apply --server-side -f https://raw.githubusercontent.com/prometheus-operator/prometheus-operator/`+version+`/example/prometheus-operator-crd/monitoring.coreos.com_servicemonitors.yaml`, FlagDry)
		sh(`kubectl apply --server-side -f https://raw.githubusercontent.com/prometheus-operator/prometheus-operator/`+version+`/example/prometheus-operator-crd/monitoring.coreos.com_thanosrulers.yaml`, FlagDry)
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
