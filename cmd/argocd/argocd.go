package argocd

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "argocd",
	Short:   "ArgoCD Utils",
	Aliases: []string{"acd"},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
