package k8s_scripts

func InstallHelloWorld(host string, dry bool) {
	sh(`helm upgrade --install \
		hello-world hello-world \
	--repo https://helm.sikalabs.io \
	--create-namespace \
	--namespace hello-world \
	--set host=`+host+` \
	--wait`, dry)
}
