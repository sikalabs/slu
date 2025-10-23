package k8s_scripts

func InstallArgoCD(namespace string, dry bool, installOnly bool) {
	helmCommand := "helm upgrade --install"
	if installOnly {
		helmCommand = "helm install"
	}
	sh(helmCommand+` \
	argocd argo-cd \
	--repo https://argoproj.github.io/argo-helm \
	--create-namespace \
	--namespace `+namespace+` \
	--wait`, dry)
}

func InstallArgoCDDomain(namespace string, domain string, dry bool, installOnly bool) {
	// https://github.com/argoproj/argo-helm/blob/main/charts/argo-cd/values.yaml
	helmCommand := "helm upgrade --install"
	if installOnly {
		helmCommand = "helm install"
	}
	sh(helmCommand+` \
	argocd argo-cd \
	--repo https://argoproj.github.io/argo-helm \
	--create-namespace \
	--namespace `+namespace+` \
	--set 'configs.cm.url'=https://`+domain+` \
	--set 'server.ingress.enabled=true' \
	--set 'server.ingress.hostname='`+domain+` \
	--set 'server.ingress.ingressClassName=nginx' \
	--set 'server.ingress.annotations.cert-manager\.io/cluster-issuer=letsencrypt' \
	--set 'server.ingress.annotations.nginx\.ingress\.kubernetes\.io/backend-protocol=HTTPS' \
	--set 'server.ingress.tls=true' \
	--wait`, dry)
}
