package vaultino

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "vaultino",
	Aliases: []string{"vt"},
	Short:   "Vaultino is custom implementation of Vault similar to Ansible Vault",
	Long:    "Vaultino is custom implementation of Vault similar to Ansible Vault. It allows you to encrypt and decrypt files using a password.",
	Example: "vaultino valutino --help",
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
