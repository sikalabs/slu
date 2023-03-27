package aws

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "aws",
	Short: "AWS Utils",
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
