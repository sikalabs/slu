package k8s_scripts

func InstallArgoCD(namespace string, dry bool) {
	sh(`helm upgrade --install \
	argocd argo-cd \
	--repo https://argoproj.github.io/argo-helm \
	--create-namespace \
	--namespace `+namespace+` \
	--wait`, dry)
}

func InstallArgoCDDomain(namespace string, domain string, dry bool) {
	sh(`helm upgrade --install \
	argocd argo-cd \
	--repo https://argoproj.github.io/argo-helm \
	--create-namespace \
	--namespace `+namespace+` \
	--set 'server.ingress.enabled=true' \
	--set 'server.ingress.hosts[0]='`+domain+` \
	--set 'server.ingress.ingressClassName=nginx' \
	--set 'server.ingress.annotations.cert-manager\.io/cluster-issuer=letsencrypt' \
	--set 'server.ingress.annotations.nginx\.ingress\.kubernetes\.io/server-snippet=proxy_ssl_verify off;' \
	--set 'server.ingress.annotations.nginx\.ingress\.kubernetes\.io/backend-protocol=HTTPS' \
	--set 'server.ingress.tls[0].hosts[0]=`+domain+`' \
	--set 'server.ingress.tls[0].secretName=argocd-tls' \
	--wait`, dry)
}
