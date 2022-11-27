package yyyy_mm_dd_hh_mm_ss

import (
	"fmt"
	"time"

	time_cmd "github.com/sikalabs/slu/cmd/time"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "yyyy-mm-dd-hh-mm-ss",
	Short: "Get time in yyyy-mm-dd-hh-mm-ss format",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		fmt.Println(time.Now().Format("2006-01-02_15-04-05"))
	},
}

func init() {
	time_cmd.Cmd.AddCommand(Cmd)
}
