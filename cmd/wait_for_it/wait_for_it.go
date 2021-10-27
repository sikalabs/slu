package wait_for_it

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var CmdFlagAddr string
var CmdFlagTimeout int

var Cmd = &cobra.Command{
	Use:     "wait-for-it",
	Short:   "Wait for TCP connection",
	Aliases: []string{"wfi"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		started := time.Now()
		for {
			_, err := net.DialTimeout("tcp", CmdFlagAddr, 100*time.Millisecond)
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
