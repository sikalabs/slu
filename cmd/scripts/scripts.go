package scripts

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "scripts",
	Short:   "Stored (Shell) Scripts",
	Aliases: []string{"s"},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
