package skip

import (
	skip_stage_cmd "github.com/sikalabs/slu/cmd/gitlab_ci/skip_stage"
	skip_stage "github.com/sikalabs/slu/utils/gitlab_ci/skip_stage"

	"github.com/spf13/cobra"
)

var CmdFlagStage string

var Cmd = &cobra.Command{
	Use:     "skip",
	Short:   "Skip stage in Gitlab CI",
	Aliases: []string{"s"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		skip_stage.SkipStage(
			skip_stage_cmd.CmdFlagConfig,
			CmdFlagStage,
		)
	},
}

func init() {
	skip_stage_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&CmdFlagStage,
		"stage",
		"s",
		"",
		"Gitlab CI Stage",
	)
	Cmd.MarkFlagRequired("stage")
}
