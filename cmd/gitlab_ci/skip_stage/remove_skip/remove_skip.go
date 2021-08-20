package remove_skip

import (
	skip_stage_cmd "github.com/sikalabs/slu/cmd/gitlab_ci/skip_stage"
	skip_stage "github.com/sikalabs/slu/utils/gitlab_ci/skip_stage"

	"github.com/spf13/cobra"
)

var CmdFlagStage string

var Cmd = &cobra.Command{
	Use:   "remove-skip",
	Short: "Remove skip stage in Gitlab CI",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		skip_stage.RemoveSkipStage(
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
