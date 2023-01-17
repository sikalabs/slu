package k8s_scripts

func InstallHelloWorld(host string, namespace string, dry bool) {
	sh(`helm upgrade --install \
		hello-world hello-world \
	--repo https://helm.sikalabs.io \
	--create-namespace \
	--namespace `+namespace+` \
	--set host=`+host+` \
	--wait`, dry)
}
