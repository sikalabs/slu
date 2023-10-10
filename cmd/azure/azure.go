package azure

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "azure",
	Aliases: []string{"az"},
	Short:   "Azure Utils",
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
