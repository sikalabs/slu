package example_server

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/sikalabs/slu/lib/file_server"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "serve-files",
	Aliases: []string{"serve"},
	Short:   "Serve files in local directory",
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		file_server.Server(8000)
	},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
