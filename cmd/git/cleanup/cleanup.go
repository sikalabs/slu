package cleanup

import (
	parent_cmd "github.com/sikalabs/slu/cmd/git"
	"github.com/sikalabs/slu/utils/git_utils"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "cleanup",
	Short: "Cleanup Git repository",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		git_utils.DeleteAllDependabotBranches()
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}
