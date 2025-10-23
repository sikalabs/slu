package k8s_scripts

func InstallIngress(useProxyProtocol bool, dry bool, installOnly bool) {
	useProxyProtocolStr := "false"
	if useProxyProtocol {
		useProxyProtocolStr = "true"
	}
	helmCommand := "helm upgrade --install"
	if installOnly {
		helmCommand = "helm install"
	}
	sh(helmCommand+` \
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

func InstallIngressAKS(
	loadBalancerIP string,
	resourceGroupName string,
	dry bool,
	installOnly bool,
) {
	helmCommand := "helm upgrade --install"
	if installOnly {
		helmCommand = "helm install"
	}
	sh(helmCommand+` \
	ingress-nginx ingress-nginx \
	--repo https://kubernetes.github.io/ingress-nginx \
	--create-namespace \
	--namespace ingress-nginx \
	--set controller.service.type=LoadBalancer \
	--set controller.ingressClassResource.default=true \
	--set controller.kind=DaemonSet \
	--set controller.hostPort.enabled=true \
	--set controller.metrics.enabled=true \
	--set controller.config.use-proxy-protocol=false \
  --set controller.service.loadBalancerIP=`+loadBalancerIP+` \
  --set controller.service.annotations.service\.beta\.kubernetes\.io/azure-load-balancer-resource-group=`+resourceGroupName+` \
  --set controller.service.annotations."service\.beta\.kubernetes\.io/azure-load-balancer-health-probe-request-path"=/healthz \
	--wait`, dry)
}
