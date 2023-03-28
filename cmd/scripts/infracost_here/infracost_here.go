package infracost_here

import (
	parent_cmd "github.com/sikalabs/slu/cmd/scripts"
	"github.com/sikalabs/slu/utils/exec_utils"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "infracost-here",
	Short:   "Run infracost breakdown in current directory",
	Aliases: []string{"ic"},
	Run: func(c *cobra.Command, args []string) {
		exec_utils.ExecOut("infracost", "breakdown", "--path", ".", "--show-skipped")
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}
