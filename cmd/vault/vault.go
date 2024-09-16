package vault

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "vault",
	Short: "HashiCorp Vault utils",
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
