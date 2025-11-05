package crypt

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "crypt",
	Short: "Encryption and decryption utilities",
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
