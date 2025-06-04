package now

import (
	"fmt"
	"strconv"
	"time"

	time_cmd "github.com/sikalabs/slu/cmd/time"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "now",
	Short: "Current time in various formats",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		now := time.Now()
		fmt.Println(strconv.Itoa(int(now.Unix())))
		fmt.Println(now.Format(time.RFC3339))
		fmt.Println(now.UTC().Format(time.RFC3339))
	},
}

func init() {
	time_cmd.Cmd.AddCommand(Cmd)
}
