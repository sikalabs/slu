package kubernetes

import (
	parent_cmd "github.com/sikalabs/slu/cmd/scripts"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "kubernetes",
	Short:   "Kubernetes Scripts",
	Aliases: []string{"k", "k8s"},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}
