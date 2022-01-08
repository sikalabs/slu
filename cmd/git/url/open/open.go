package open

import (
	"fmt"

	"github.com/pkg/browser"
	parent_cmd "github.com/sikalabs/slu/cmd/git/url"
	"github.com/sikalabs/slu/utils/git_utils"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "open",
	Short: "Open repository in browser",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		url := git_utils.GetRepoUrl()
		fmt.Println(url)
		browser.OpenURL(url)
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}
