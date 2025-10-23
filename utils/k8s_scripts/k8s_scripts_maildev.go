package k8s_scripts

func InstallMaildevDomain(namespace string, domain string, dry bool, installOnly bool) {
	helmCommand := "helm upgrade --install"
	if installOnly {
		helmCommand = "helm install"
	}
	sh(helmCommand+` \
	maildev maildev \
	--repo https://helm.sikalabs.io \
	--create-namespace \
	--namespace `+namespace+` \
	--set host=`+domain+` \
	--wait`, dry)
}
