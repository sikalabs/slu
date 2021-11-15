package unix

import (
	"fmt"
	"strconv"
	"time"

	time_cmd "github.com/sikalabs/slu/cmd/time"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "unix",
	Short: "Get Unix time",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		fmt.Println(strconv.Itoa(int(time.Now().Unix())))
	},
}

func init() {
	time_cmd.Cmd.AddCommand(Cmd)
}
