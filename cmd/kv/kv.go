package k8s

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "kv",
	Short: "Key-Value store in Vault",
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
