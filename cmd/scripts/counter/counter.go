package download

import (
	"fmt"
	"time"

	parent_cmd "github.com/sikalabs/slu/cmd/scripts"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "counter",
	Short: "Count 1, 2, 3, ...",
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
