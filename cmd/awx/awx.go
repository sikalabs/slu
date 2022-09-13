package awx

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "awx",
	Short: "AWX Utils",
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
