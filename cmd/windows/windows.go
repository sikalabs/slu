package argocd

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "windows",
	Short:   "Windows Utils",
	Aliases: []string{"win"},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
