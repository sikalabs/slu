package delete_all_local_branches

import (
	parent_cmd "github.com/sikalabs/slu/cmd/git"
	"github.com/sikalabs/slu/utils/git_utils"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "delete-all-local-branches",
	Short:   "Delete all local brancher (instead of current)",
	Aliases: []string{"dalb"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		git_utils.DeleteAllLocalBranches()
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}
