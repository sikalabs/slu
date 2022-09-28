package k8s_scripts

func InstallIngress(useProxyProtocol bool, dry bool) {
	useProxyProtocolStr := "false"
	if useProxyProtocol {
		useProxyProtocolStr = "true"
	}
	sh(`helm upgrade --install \
	ingress-nginx ingress-nginx \
	--repo https://kubernetes.github.io/ingress-nginx \
	--create-namespace \
	--namespace ingress-nginx \
	--set controller.service.type=ClusterIP \
	--set controller.ingressClassResource.default=true \
	--set controller.kind=DaemonSet \
	--set controller.hostPort.enabled=true \
	--set controller.metrics.enabled=true \
	--set controller.config.use-proxy-protocol=`+useProxyProtocolStr+` \
	--wait`, dry)
}
