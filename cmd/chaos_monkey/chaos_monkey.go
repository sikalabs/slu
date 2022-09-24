package chaos_monkey

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "chaos-monkey",
	Short:   "Chaos Monkey Utils",
	Aliases: []string{"cm"},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
