package k8s

import "github.com/sikalabs/slu/utils/exec_utils"

func AddToKubeconfigShell(path string) error {
	return exec_utils.ExecOut("sh", "-c", `
	cp ~/.kube/config ~/.kube/.config.$(date +%Y-%m-%d_%H-%M-%S).backup
	KUBECONFIG=`+path+`:~/.kube/config kubectl config view --raw > /tmp/kubeconfig.merge.yml && cp /tmp/kubeconfig.merge.yml ~/.kube/config
	chmod 600 ~/.kube/config
	`)
}
