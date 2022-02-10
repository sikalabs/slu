package watch

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var FlagSleepTime int
var FlagVerbose bool

var Cmd = &cobra.Command{
	Use:     "watch",
	Short:   "Watch Util",
	Aliases: []string{"w"},
	Run: func(c *cobra.Command, args []string) {
		i := 0
		for {
			strArgs := strings.Join(args, " ")
			if FlagVerbose {
				fmt.Printf("%d: %s\n", i, strArgs)
			}
			cmd := exec.Command("/bin/sh", "-c", strArgs)
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				if FlagVerbose {
					log.Println(err)
				}
			}
			time.Sleep(time.Duration(FlagSleepTime) * time.Millisecond)
			i++
		}
	},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
	Cmd.Flags().IntVarP(
		&FlagSleepTime,
		"sleep-time",
		"s",
		300,
		"Sleep time in ms",
	)
	Cmd.Flags().BoolVarP(
		&FlagVerbose,
		"verbose",
		"v",
		false,
		"Verbose output",
	)
}
