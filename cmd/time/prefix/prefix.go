package unix

import (
	"bufio"
	"fmt"
	"os"
	"time"

	time_cmd "github.com/sikalabs/slu/cmd/time"
	"github.com/sikalabs/slu/utils/time_utils"

	"github.com/spf13/cobra"
)

var CmdFlagNoSince bool
var CmdFlagNoTimestamp bool
var CmdFlagSeparator bool

var Cmd = &cobra.Command{
	Use:   "prefix",
	Short: "Prefix stdin with timestap",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		start := time.Now()
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			since := time.Since(start)
			if !CmdFlagNoTimestamp {
				fmt.Printf("%s ", time.Now().Format("2006-01-02 15:04:05"))
			}
			if !CmdFlagNoSince {
				fmt.Printf("%s ", time_utils.DurationToString(since))
			}
			if CmdFlagSeparator {
				fmt.Print("| ")
			}
			fmt.Println(s.Text())
		}
	},
}

func init() {
	time_cmd.Cmd.AddCommand(Cmd)
	Cmd.PersistentFlags().BoolVarP(
		&CmdFlagNoSince,
		"no-since",
		"S",
		false,
		"No since time (from start)",
	)
	Cmd.PersistentFlags().BoolVarP(
		&CmdFlagNoTimestamp,
		"no-timestamp",
		"T",
		false,
		"No timestamp",
	)
	Cmd.PersistentFlags().BoolVarP(
		&CmdFlagSeparator,
		"separator",
		"s",
		false,
		"Add separator (\" | \")",
	)
}
