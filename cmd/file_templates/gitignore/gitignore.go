package gitignore

import (
	file_templates_cmd "github.com/sikalabs/slu/cmd/file_templates"
	"github.com/sikalabs/slu/file_templates/gitignore"

	"github.com/spf13/cobra"
)

var FlagTerraform bool

var Cmd = &cobra.Command{
	Use:     "gitignore",
	Short:   "Create basic gitignore",
	Aliases: []string{"gi"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		content := gitignore.GitignoreBase
		if FlagTerraform {
			content += "\n" + gitignore.GitignoreTerraform
		}
		gitignore.CreateGitignore(content)
	},
}

func init() {
	file_templates_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().BoolVar(
		&FlagTerraform,
		"terraform",
		false,
		"Add Terraform part",
	)
}
