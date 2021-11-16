package gitlab

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "gitlab",
	Short:   "Utils for Gitlab",
	Aliases: []string{"gl"},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
