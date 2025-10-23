package k8s_scripts

func InstallCNPG(dry bool, installOnly bool) {
	helmCommand := "helm upgrade --install"
	if installOnly {
		helmCommand = "helm install"
	}
	sh(helmCommand+` \
cnpg cloudnative-pg \
--repo https://cloudnative-pg.github.io/charts \
--namespace cnpg-system \
--create-namespace \
--wait`, dry)
}
