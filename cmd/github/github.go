package github

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "github",
	Short:   "Github Utils",
	Aliases: []string{"gh"},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
