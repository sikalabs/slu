package k8s_scripts

func InstallCNPG(dry bool) {
	sh(`helm upgrade --install \
cnpg cloudnative-pg \
--repo https://cloudnative-pg.github.io/charts \
--namespace cnpg-system \
--create-namespace \
--wait`, dry)
}
