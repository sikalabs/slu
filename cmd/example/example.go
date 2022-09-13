package example

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "example",
	Short: "Example commands",
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
