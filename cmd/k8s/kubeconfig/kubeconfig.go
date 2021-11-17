package kubeconfig

import (
	k8s_cmd "github.com/sikalabs/slu/cmd/k8s"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "kubeconfig",
	Short:   "kubeconfig utils",
	Aliases: []string{"config", "conf"},
}

func init() {
	k8s_cmd.Cmd.AddCommand(Cmd)
}
