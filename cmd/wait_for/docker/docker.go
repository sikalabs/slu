package docker

import (
	parentcmd "github.com/sikalabs/slu/cmd/wait_for"
	"github.com/sikalabs/slu/utils/wait_for_docker_utils"
	"github.com/spf13/cobra"
)

var FlagTimeout int

var Cmd = &cobra.Command{
	Use:   "docker",
	Short: "Wait for Docker",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		wait_for_docker_utils.WaitForDocker(FlagTimeout)
	},
}

func init() {
	parentcmd.Cmd.AddCommand(Cmd)
	Cmd.MarkFlagRequired("address")
	Cmd.Flags().IntVarP(
		&FlagTimeout,
		"timeout",
		"t",
		5*60, // 5 min
		"Timeout",
	)
}
