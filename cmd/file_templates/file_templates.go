package file_templates

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "file-templates",
	Short:   "Create common files from templates",
	Aliases: []string{"ft"},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
