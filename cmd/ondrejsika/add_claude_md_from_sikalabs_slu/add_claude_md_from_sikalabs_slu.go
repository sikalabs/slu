package add_claude_md_from_sikalabs_slu

import (
	"github.com/sikalabs/slu/cmd/ondrejsika"
	"github.com/sikalabs/slu/utils/exec_utils"
	"github.com/spf13/cobra"
)

func init() {
	ondrejsika.Cmd.AddCommand(Cmd)
}

var Cmd = &cobra.Command{
	Use:     "add-claude-md-from-sikalabs-slu",
	Short:   "Add CLAUDE.md from sikalabs/slu repository",
	Aliases: []string{"add-claude-md"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		addClaudeMdFromSikalasbSlu()
	},
}

func addClaudeMdFromSikalasbSlu() {
	exec_utils.ExecOut("curl", "-fsSL", "https://raw.githubusercontent.com/sikalabs/slu/refs/heads/master/CLAUDE.md", "-o", "CLAUDE.md")
	exec_utils.ExecOut("git", "add", "CLAUDE.md")
	exec_utils.ExecOut("git", "commit", "-m", "feat(CLAUDE): Add CLAUDE.md from sikalabs/slu")
}
