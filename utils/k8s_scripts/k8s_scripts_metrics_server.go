package k8s_scripts

func InstallMetricsServer(dry bool) {
	sh(`kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml`, dry)
}
