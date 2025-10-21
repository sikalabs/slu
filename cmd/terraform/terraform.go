package terraform

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "terraform",
	Short:   "Terraform Utils",
	Aliases: []string{"tf"},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
