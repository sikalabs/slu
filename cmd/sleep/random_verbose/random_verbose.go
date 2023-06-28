package random_verbose

import (
	"fmt"
	"math/rand"
	"time"

	sleep_cmd "github.com/sikalabs/slu/cmd/sleep"
	"github.com/spf13/cobra"
)

var FlagMinTime int
var FlagMaxTime int

var Cmd = &cobra.Command{
	Use:   "random-verbose",
	Short: "Sleep random time with verbose output",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		rand.Seed(time.Now().UnixNano())

		sleepTimeInSeconds := rand.Intn(FlagMaxTime-FlagMinTime) + FlagMinTime

		fmt.Printf("Sleep %d seconds\n", sleepTimeInSeconds)
		for i := 0; i < sleepTimeInSeconds; i++ {
			time.Sleep(time.Second)
			fmt.Printf("... %d/%d seconds\n", i, sleepTimeInSeconds)
		}
		fmt.Println("Done.")
	},
}

func init() {
	sleep_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().IntVar(
		&FlagMinTime,
		"min",
		0,
		"Minimum sleep time (in seconds)",
	)
	Cmd.Flags().IntVar(
		&FlagMaxTime,
		"max",
		10, // 10s
		"Maximum sleep time (in seconds)",
	)
}
