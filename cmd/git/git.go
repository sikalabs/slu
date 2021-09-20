package git

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "git",
	Short:   "Git Utils",
	Aliases: []string{"g"},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
