package tag_next_calver

import (
	parent_cmd "github.com/sikalabs/slu/cmd/git"
	"github.com/sikalabs/slu/utils/git_utils"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "tag-next-calver",
	Short: "Crate tag for next CalVer (vYYYY.MM.MICRO)",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		git_utils.TagNextCalver()
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}
