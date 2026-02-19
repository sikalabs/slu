package k8s_scripts

func InstallKafdropOpenshift(name, namespace, host, kafkaBootstrap string, dry bool) {
	sh(`helm upgrade --install \
	`+name+` \
	--namespace `+namespace+` \
	--create-namespace \
	--repo https://helm.sikalabs.io \
	simple-kafdrop \
	--set host=`+host+` \
	--set kafkaBootstrapServer=`+kafkaBootstrap+` \
	--set tls=false \
	--set ingressClassName=openshift-default \
	--set ingressExtraAnnotations."route\.openshift\.io/termination"=edge \
	--set ingressExtraAnnotations."route\.openshift\.io/insecureEdgeTerminationPolicy"=Redirect \
	--wait`, dry)
}
