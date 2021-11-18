package wait_for_tls

import (
	"crypto/tls"
	"fmt"
	"os"
	"time"

	"github.com/sikalabs/slu/cmd/root"
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
		started := time.Now()
		for {
			_, err := tls.Dial("tcp", CmdFlagAddr, &tls.Config{
				InsecureSkipVerify: false,
			})
			if err == nil {
				os.Exit(0)
			}
			fmt.Println(err)
			if time.Since(started) > time.Duration(CmdFlagTimeout*int(time.Second)) {
				os.Exit(1)
			}
			time.Sleep(100 * time.Millisecond)
		}
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
