package wait_for_tls

import (
	"crypto/tls"
	"time"

	"github.com/sikalabs/slu/cmd/root"
	"github.com/sikalabs/slu/utils/wait_for_utils"
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
		wait_for_utils.WaitFor(
			CmdFlagTimeout, 100*time.Millisecond,
			func() (bool, bool, string, error) {
				_, err := tls.Dial("tcp", CmdFlagAddr, &tls.Config{
					InsecureSkipVerify: false,
				})
				if err == nil {
					return wait_for_utils.WaitForResponseSucceeded("TLS certificate validated")
				}
				return wait_for_utils.WaitForResponseWaiting(err.Error())
			},
		)
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
