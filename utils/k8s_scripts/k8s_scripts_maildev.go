package k8s_scripts

func InstallMaildevDomain(namespace string, domain string, dry bool) {
	sh(`helm upgrade --install \
	maildev maildev \
	--repo https://helm.sikalabs.io \
	--create-namespace \
	--namespace `+namespace+` \
	--set host=`+domain+` \
	--wait`, dry)
}
