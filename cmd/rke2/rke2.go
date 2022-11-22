package rke2

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "rke2",
	Short: "RKE2 Utils",
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
