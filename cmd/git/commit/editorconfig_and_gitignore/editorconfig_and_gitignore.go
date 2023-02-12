package editorconfig_and_gitignore

import (
	commit_cmd "github.com/sikalabs/slu/cmd/git/commit"
	"github.com/sikalabs/slu/utils/exec_utils"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "editorconfig-and-gitignore",
	Short:   "Commit change: Create .editorconfig and .gitignore",
	Aliases: []string{},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		exec_utils.ExecOut("git", "add", ".editorconfig", ".gitignore")
		exec_utils.ExecOut(
			"git", "commit", "-n", "-m",
			"chore: Create .editorconfig and .gitignore")
	},
}

func init() {
	commit_cmd.Cmd.AddCommand(Cmd)
}
