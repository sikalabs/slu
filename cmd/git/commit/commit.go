package commit

import (
	git_cmd "github.com/sikalabs/slu/cmd/git"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "commit",
	Short:   "Commit specific changes",
	Aliases: []string{"c"},
}

func init() {
	git_cmd.Cmd.AddCommand(Cmd)
}
