package golang

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "golang",
	Short:   "Go Tools",
	Aliases: []string{"go"},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
