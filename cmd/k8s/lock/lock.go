package lock

import (
	k8s_cmd "github.com/sikalabs/slu/cmd/k8s"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "lock",
	Short:   "Locks in Kubernetes",
	Aliases: []string{"l"},
}

func init() {
	k8s_cmd.Cmd.AddCommand(Cmd)
}
