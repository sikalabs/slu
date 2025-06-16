package mon

import (
	"log"
	"time"

	parent_cmd "github.com/sikalabs/slu/cmd/tools"
	"github.com/sikalabs/slu/pkg/tools/mon"
	"github.com/spf13/cobra"
)

var FlagSleepTime int

var Cmd = &cobra.Command{
	Use:   "mon",
	Short: "Simple VM monitoring tool",
	Run: func(c *cobra.Command, args []string) {
		for {
			mon.Mon()
			if FlagSleepTime <= 0 {
				log.Println("Invalid sleep time, using default of 10 seconds")
				time.Sleep(10 * time.Second)
			} else {
				log.Printf("Sleeping for %d minutes", FlagSleepTime)
				time.Sleep(time.Duration(FlagSleepTime) * time.Minute)
			}
		}
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().IntVarP(
		&FlagSleepTime,
		"sleep",
		"s",
		10,
		"Sleep time in minutes between checks",
	)
}
