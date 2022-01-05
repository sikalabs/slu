package generate_files

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "generate-files",
	Short:   "Generate dummy files for testing",
	Aliases: []string{"gen-files"},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
