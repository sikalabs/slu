package gitlab_ci

import (
	parent_cmd "github.com/sikalabs/slu/cmd/scripts"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "gitlab-ci",
	Short:   "Gitlab CI Scripts",
	Aliases: []string{"gci"},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}
