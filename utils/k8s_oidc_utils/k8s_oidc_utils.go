package k8s_oidc_utils

import (
	"fmt"

	"github.com/sikalabs/slu/utils/exec_utils"
)

func CreateOidcUser(
	name, issuerUrl, clientId, clientSecret string,
	dry bool,
) {
	args := []string{
		"config", "set-credentials", name,
		"--exec-api-version=client.authentication.k8s.io/v1beta1",
		"--exec-command=kubectl",
		"--exec-arg=oidc-login",
		"--exec-arg=get-token",
		"--exec-arg=--oidc-issuer-url=" + issuerUrl,
		"--exec-arg=--oidc-client-id=" + clientId,
	}
	if clientSecret != "" {
		args = append(args, "--exec-arg=--oidc-client-secret="+clientSecret)
	}

	if dry {
		fmt.Print("kubectl")
		for _, arg := range args {
			fmt.Printf(" %s", arg)
		}
		fmt.Println("")
	} else {
		exec_utils.ExecOut("kubectl", args...)
	}
}
