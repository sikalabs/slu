package scripts

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "scripts",
	Short:   "Go & Shell Scripts",
	Aliases: []string{"script", "s"},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
