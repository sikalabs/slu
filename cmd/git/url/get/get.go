package get

import (
	"fmt"

	parent_cmd "github.com/sikalabs/slu/cmd/git/url"
	"github.com/sikalabs/slu/utils/git_utils"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "get",
	Short: "Get browser URL of repository",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		fmt.Println(git_utils.GetRepoUrl())
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}
