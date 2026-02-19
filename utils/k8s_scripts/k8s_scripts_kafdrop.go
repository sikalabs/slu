package k8s_scripts

func InstallKafdrop(name, namespace, host, kafkaBootstrap string, dry bool) {
	sh(`helm upgrade --install \
	`+name+` \
	--namespace `+namespace+` \
	--create-namespace \
	--repo https://helm.sikalabs.io \
	simple-kafdrop \
	--set host=`+host+` \
	--set kafkaBootstrapServer=`+kafkaBootstrap+` \
	--wait`, dry)
}
