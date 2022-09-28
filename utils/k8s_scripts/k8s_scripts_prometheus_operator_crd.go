package k8s_scripts

import "github.com/sikalabs/slu/utils/github_utils"

func InstallPrometheusOperatorCRD(version string, dry bool) {
	if version == "latest" {
		version = github_utils.GetLatestRelease("prometheus-operator", "prometheus-operator")
	}
	sh(`kubectl apply --server-side -f https://raw.githubusercontent.com/prometheus-operator/prometheus-operator/`+version+`/example/prometheus-operator-crd/monitoring.coreos.com_alertmanagerconfigs.yaml`, dry)
	sh(`kubectl apply --server-side -f https://raw.githubusercontent.com/prometheus-operator/prometheus-operator/`+version+`/example/prometheus-operator-crd/monitoring.coreos.com_alertmanagers.yaml`, dry)
	sh(`kubectl apply --server-side -f https://raw.githubusercontent.com/prometheus-operator/prometheus-operator/`+version+`/example/prometheus-operator-crd/monitoring.coreos.com_podmonitors.yaml`, dry)
	sh(`kubectl apply --server-side -f https://raw.githubusercontent.com/prometheus-operator/prometheus-operator/`+version+`/example/prometheus-operator-crd/monitoring.coreos.com_probes.yaml`, dry)
	sh(`kubectl apply --server-side -f https://raw.githubusercontent.com/prometheus-operator/prometheus-operator/`+version+`/example/prometheus-operator-crd/monitoring.coreos.com_prometheuses.yaml`, dry)
	sh(`kubectl apply --server-side -f https://raw.githubusercontent.com/prometheus-operator/prometheus-operator/`+version+`/example/prometheus-operator-crd/monitoring.coreos.com_prometheusrules.yaml`, dry)
	sh(`kubectl apply --server-side -f https://raw.githubusercontent.com/prometheus-operator/prometheus-operator/`+version+`/example/prometheus-operator-crd/monitoring.coreos.com_servicemonitors.yaml`, dry)
	sh(`kubectl apply --server-side -f https://raw.githubusercontent.com/prometheus-operator/prometheus-operator/`+version+`/example/prometheus-operator-crd/monitoring.coreos.com_thanosrulers.yaml`, dry)
}
