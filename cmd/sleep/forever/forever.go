package forever

import (
	"time"

	sleep_cmd "github.com/sikalabs/slu/cmd/sleep"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "forever",
	Short:   "Sleep forever",
	Aliases: []string{"f", "inf"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		for {
			time.Sleep(time.Second)
		}
	},
}

func init() {
	sleep_cmd.Cmd.AddCommand(Cmd)
}
