package mit_license

import (
	commit_cmd "github.com/sikalabs/slu/cmd/git/commit"
	"github.com/sikalabs/slu/utils/exec_utils"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "mit-license",
	Aliases: []string{"mit"},
	Short:   "Commit change: Create MIT LICENSE file",
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		exec_utils.ExecOut("git", "add", "LICENSE")
		exec_utils.ExecOut(
			"git", "commit", "-n", "-m",
			"feat(LICENSE): Add MIT LICENSE")
	},
}

func init() {
	commit_cmd.Cmd.AddCommand(Cmd)
}
