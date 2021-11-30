package wait_for

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "wait-for",
	Short:   "Wait for utils",
	Aliases: []string{"wf"},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
