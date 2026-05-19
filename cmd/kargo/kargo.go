package kargo

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "kargo",
	Short:   "Kargo.io Utils",
	Aliases: []string{"kg"},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
