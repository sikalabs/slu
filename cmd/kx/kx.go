package kx

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/sikalabs/slu/utils/exec_utils"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "kx [<context>]",
	Short: "Simple embdeded kubectx",
	Args:  cobra.MaximumNArgs(1),
	Run: func(c *cobra.Command, args []string) {
		if len(args) == 0 {
			exec_utils.ExecOut("kubectl", "config", "get-contexts", "-o", "name")
		} else {
			exec_utils.ExecOut("kubectl", "config", "use-context", args[0])
		}
	},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
