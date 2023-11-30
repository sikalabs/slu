package web_server

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "web-server",
	Short: "Set of embedded web servers",
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
