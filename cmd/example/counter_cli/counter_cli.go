package counter_cli

import (
	"fmt"
	"time"

	parent_cmd "github.com/sikalabs/slu/cmd/example"
	"github.com/spf13/cobra"
)

var FlagPort int

var Cmd = &cobra.Command{
	Use:   "counter-cli",
	Short: "Example CLI counter",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		i := 0
		for {
			fmt.Println(i)
			i++
			time.Sleep(1 * time.Second)
		}
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}
