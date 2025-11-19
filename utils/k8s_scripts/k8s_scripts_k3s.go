package k8s_scripts

func InstallK3s(dry bool) {
	sh(`curl -sfL https://get.k3s.io | sh -s - --disable traefik`, dry)
}
