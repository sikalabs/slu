package k8s_scripts

func InstallFlux(dry bool) {
	sh(`helm install \
-n flux-system \
--create-namespace \
flux \
oci://ghcr.io/fluxcd-community/charts/flux2`, dry)
}
