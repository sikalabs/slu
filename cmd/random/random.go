package random

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "random",
	Short:   "Random Utils",
	Aliases: []string{"rnd", "r"},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
