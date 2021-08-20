package gitlab_ci

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "gitlab-ci",
	Short:   "Utils for Gitlab CI",
	Aliases: []string{"gci"},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
