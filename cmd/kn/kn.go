package kn

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/sikalabs/slu/utils/exec_utils"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "kn [<namespace>]",
	Short: "Simple embdeded kubens",
	Args:  cobra.MaximumNArgs(1),
	Run: func(c *cobra.Command, args []string) {
		if len(args) == 0 {
			exec_utils.ExecOut("kubectl", "get", "namespaces", "-o", `jsonpath={range .items[*]}{.metadata.name}{"\n"}{end}`)
		} else {
			exec_utils.ExecOut("kubectl", "config", "set-context", "--current", "--namespace", args[0])
		}
	},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
