package run_redis

import (
	parent_cmd "github.com/sikalabs/slu/cmd/scripts/docker"
	"github.com/sikalabs/slu/utils/exec_utils"
	"github.com/spf13/cobra"
)

var FlagDry bool
var FlagNamespace string
var FlagDomain string
var FlagInstallOnly bool
var FlagPassword string

var Cmd = &cobra.Command{
	Use:     "run-redis",
	Short:   "docker run --name redis -d -p 6379:6379 redis",
	Aliases: []string{"rr"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		dockerArgs := []string{
			"run",
			"--name", "redis",
			"-d",
			"-p", "6379:6379",
		}

		if FlagPassword != "" {
			dockerArgs = append(dockerArgs, "redis", "redis-server", "--requirepass", FlagPassword)
		} else {
			dockerArgs = append(dockerArgs, "redis")
		}

		exec_utils.ExecOut("docker", dockerArgs...)
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagPassword,
		"password",
		"p",
		"",
		"Set Redis password (requirepass)",
	)
}
