package skip_stage

import (
	gitlab_ci_cmd "github.com/sikalabs/slu/cmd/gitlab_ci"

	"github.com/spf13/cobra"
)

var CmdFlagConfig string

var Cmd = &cobra.Command{
	Use:     "skip-stage",
	Short:   "Skip stage in Gitlab CI",
	Aliases: []string{"ss"},
}

func init() {
	gitlab_ci_cmd.Cmd.AddCommand(Cmd)
	Cmd.PersistentFlags().StringVarP(
		&CmdFlagConfig,
		"config",
		"c",
		"",
		"skip-stage config (JSON)",
	)
	Cmd.MarkPersistentFlagRequired("config")
}
