package setup_rke2_single_node

import (
	"fmt"
	"log"
	"runtime"

	parent_cmd "github.com/sikalabs/slu/cmd/scripts"
	"github.com/sikalabs/slu/utils/random_utils"
	"github.com/sikalabs/slu/utils/sh_utils"
	"github.com/spf13/cobra"
)

var FlagDry bool
var FlagTlsSan string

var Cmd = &cobra.Command{
	Use:   "setup-rke2-single-node",
	Short: "Setup RKE2 Single Node",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		if runtime.GOOS != "linux" && FlagDry != true {
			log.Fatalln("You can't setup RKE cluser on " + runtime.GOOS +
				". Try with --dry .")
		}
		sh("slu install-bin kubectl", FlagDry)
		sh("slu install-bin helm", FlagDry)
		sh("slu install-bin k9s", FlagDry)
		sh(`curl -sfL https://get.rke2.io | sh -
mkdir -p /etc/rancher/rke2/
cat << EOF > /etc/rancher/rke2/config.yaml
token: `+getPasswordOrDie()+`
tls-san:
- `+FlagTlsSan+`
disable:
- rke2-ingress-nginx
EOF
systemctl enable rke2-server.service
systemctl start rke2-server.service
`, FlagDry)
		sh("mkdir -p ~/.kube", FlagDry)
		sh("cp /etc/rancher/rke2/rke2.yaml ~/.kube/config", FlagDry)
		sh("echo >> ~/.bashrc", FlagDry)
		sh("echo 'source <(kubectl completion bash)' >> ~/.bashrc", FlagDry)
		sh("echo 'alias k=kubectl' >> ~/.bashrc", FlagDry)
		sh("echo 'complete -o default -F __start_kubectl k' >> ~/.bashrc", FlagDry)
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().BoolVar(
		&FlagDry,
		"dry",
		false,
		"Dry run",
	)
	Cmd.Flags().StringVar(
		&FlagTlsSan,
		"tls-san",
		"",
		"TLS SAN domain (like rke2.sikademo.com)",
	)
	Cmd.MarkFlagRequired("tls-san")
}

func sh(script string, dry bool) {
	if dry {
		fmt.Println(script)
		return
	}
	err := sh_utils.ExecShOutDir("", script)
	if err != nil {
		sh_utils.HandleError(err)
	}
}

func getPasswordOrDie() string {
	password, err := random_utils.RandomPassword()
	if err != nil {
		log.Fatal(err)
	}
	return password
}
