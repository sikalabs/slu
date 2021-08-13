package editorconfig

import (
	file_templates_cmd "github.com/sikalabs/slu/cmd/file_templates"
	"github.com/sikalabs/slu/file_templates/editorconfig"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "editorconfig",
	Short:   "Create basic editorconfig",
	Aliases: []string{"ec"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		editorconfig.CreateEditorconfig()
	},
}

func init() {
	file_templates_cmd.Cmd.AddCommand(Cmd)
}
