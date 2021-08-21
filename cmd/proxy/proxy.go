package proxy

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "proxy",
	Short: "Proxy Utils",
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
