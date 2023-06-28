package k8s

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "k8s",
	Short:   "Utils for Kubernetes",
	Aliases: []string{"k", "kubernetes"},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
