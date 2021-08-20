package show

import (
	skip_stage_cmd "github.com/sikalabs/slu/cmd/gitlab_ci/skip_stage"
	skip_stage "github.com/sikalabs/slu/utils/gitlab_ci/skip_stage"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "show",
	Short: "Show skip-stage variables in Gitlab CI",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		skip_stage.SkipStageShow(
			skip_stage_cmd.CmdFlagConfig,
		)
	},
}

func init() {
	skip_stage_cmd.Cmd.AddCommand(Cmd)
}
