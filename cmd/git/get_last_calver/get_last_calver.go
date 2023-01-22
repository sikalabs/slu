package get_last_calver

import (
	parent_cmd "github.com/sikalabs/slu/cmd/git"
	"github.com/sikalabs/slu/utils/git_utils"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "get-last-calver",
	Short: "Get last CalVer (vYYYY.MM.MICRO) tag if exists",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		git_utils.PrintLastCalverIfExistsOrFail()
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}
