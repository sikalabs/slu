package file_utils

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "file-utils",
	Short:   "Utils for working with files like sed, ...",
	Aliases: []string{"fu"},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
