package shell_scripts

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "shell-scripts",
	Short:   "Shell Scripts Utils",
	Aliases: []string{"sh"},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
