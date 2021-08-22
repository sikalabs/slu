package k8s

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "tls",
	Short: "TLS Utils",
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
