package random

import (
	"math/rand"
	"time"

	sleep_cmd "github.com/sikalabs/slu/cmd/sleep"
	"github.com/spf13/cobra"
)

var FlagMinTime int
var FlagMaxTime int

var Cmd = &cobra.Command{
	Use:   "random",
	Short: "Sleep random time",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Intn(FlagMaxTime-FlagMinTime)+FlagMinTime) * time.Millisecond)
	},
}

func init() {
	sleep_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().IntVar(
		&FlagMinTime,
		"min",
		0,
		"Minimum sleep time (in ms)",
	)
	Cmd.Flags().IntVar(
		&FlagMaxTime,
		"max",
		1000, // 1s
		"Maximum sleep time (in ms)",
	)
}
