package mail

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "mail",
	Short:   "Mail Utils",
	Aliases: []string{"m"},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
