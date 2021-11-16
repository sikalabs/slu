package prune_environments

import (
	gitlab_cmd "github.com/sikalabs/slu/cmd/gitlab"
	"github.com/sikalabs/slu/utils/gitlab_utils/prune_environments"

	"github.com/spf13/cobra"
)

var CmdFlagConfig string

var Cmd = &cobra.Command{
	Use:   "prune-environments",
	Short: "Prune (stop & remove) all project environments",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		prune_environments.PruneEnvironments(CmdFlagConfig)
	},
}

func init() {
	gitlab_cmd.Cmd.AddCommand(Cmd)
	Cmd.PersistentFlags().StringVarP(
		&CmdFlagConfig,
		"config",
		"c",
		"",
		"skip-stage config (JSON)",
	)
	Cmd.MarkPersistentFlagRequired("config")
}
