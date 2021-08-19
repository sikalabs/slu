package get

import (
	k8s_cmd "github.com/sikalabs/slu/cmd/k8s"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "get",
	Short: "Get data from Secret or ConfigMap",
}

func init() {
	k8s_cmd.Cmd.AddCommand(Cmd)
}
