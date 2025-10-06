package unix_from_year

import (
	"fmt"
	"strconv"
	"time"

	time_cmd "github.com/sikalabs/slu/cmd/time"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "unix-from-year [year]",
	Short: "Get Unix timestamp for January 1st of the given year",
	Args:  cobra.ExactArgs(1),
	Run: func(c *cobra.Command, args []string) {
		year, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Error: Invalid year")
			return
		}
		t := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
		fmt.Println(t.Unix())
	},
}

func init() {
	time_cmd.Cmd.AddCommand(Cmd)
}
