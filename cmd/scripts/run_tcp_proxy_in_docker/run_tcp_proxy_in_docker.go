package run_tcp_proxy_in_docker

import (
	parent_cmd "github.com/sikalabs/slu/cmd/scripts"
	"github.com/sikalabs/slu/utils/exec_utils"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "run-tcp-proxy-in-docker <external_port> <internal_port>",
	Short:   "Run slu tcp proxy in Docker",
	Args:    cobra.ExactArgs(2),
	Aliases: []string{"run-proxy"},
	Run: func(c *cobra.Command, args []string) {
		e := args[0]
		i := args[1]
		exec_utils.ExecOut(
			"docker", "run", "-d", "--name", "slu-tcp-proxy-"+e+"-"+i,
			"--network", "host", "sikalabs/slu:v0.60.0",
			"slu", "proxy", "tcp", "-l", ":"+e, "-r", "127.0.0.1:"+i,
		)
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}
