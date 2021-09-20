package pkg_if

import (
	git_cmd "github.com/sikalabs/slu/cmd/git"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "if",
	Short: "Git IF Utils",
}

func init() {
	git_cmd.Cmd.AddCommand(Cmd)
}
