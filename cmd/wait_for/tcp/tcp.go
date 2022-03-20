package tcp

import (
	parentcmd "github.com/sikalabs/slu/cmd/wait_for"
	"github.com/sikalabs/slu/utils/wait_for_tcp_utils"
	"github.com/spf13/cobra"
)

var CmdFlagAddr string
var CmdFlagTimeout int

var Cmd = &cobra.Command{
	Use:   "tcp",
	Short: "Wait for TCP connection",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		wait_for_tcp_utils.WaitForTcp(CmdFlagTimeout, CmdFlagAddr)
	},
}

func init() {
	parentcmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&CmdFlagAddr,
		"address",
		"a",
		"",
		"Address (eg.: google.com:443)",
	)
	Cmd.MarkFlagRequired("address")
	Cmd.Flags().IntVarP(
		&CmdFlagTimeout,
		"timeout",
		"t",
		5*60, // 5 min
		"Timeout",
	)
}
