package tools

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "tools",
	Short: "Tools build into slu",
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
