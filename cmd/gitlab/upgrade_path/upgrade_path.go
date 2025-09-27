package upgrade_path

import (
	"fmt"

	parent_cmd "github.com/sikalabs/slu/cmd/gitlab"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "upgrade-path",
	Short: "Get Gitlab Upgrade Path URL",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(GetGitlabUpgradePathURL())
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}
