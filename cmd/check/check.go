package check

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "check",
	Short: "Check for dependencies, versions, clusters, ...",
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
