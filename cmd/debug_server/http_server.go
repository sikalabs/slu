package debug_server

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "debug-server",
	Short:   "HTTP Server for debug & development",
	Aliases: []string{"dserver", "dese"},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
