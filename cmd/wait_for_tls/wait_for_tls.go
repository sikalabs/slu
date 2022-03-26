package wait_for_tls

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/sikalabs/slu/utils/wait_for_tls_utils"
	"github.com/spf13/cobra"
)

var CmdFlagAddr string
var CmdFlagTimeout int

var Cmd = &cobra.Command{
	Use:     "wait-for-tls",
	Short:   "Wait for VALID TLS certificate",
	Aliases: []string{"wftls"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		wait_for_tls_utils.WaitForTls(CmdFlagTimeout, CmdFlagAddr)
	},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
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
