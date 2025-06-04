package parse

import (
	"fmt"
	"log"
	"strconv"
	"time"

	time_cmd "github.com/sikalabs/slu/cmd/time"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "parse <unix>",
	Short: "Parse time to various formats",
	Args:  cobra.ExactArgs(1),
	Run: func(c *cobra.Command, args []string) {
		unixTimeInt, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("Error parsing unix time: %v\n", err)
		}
		t := time.Unix(int64(unixTimeInt), 0)
		fmt.Println(t.Format(time.RFC3339))
		fmt.Println(t.UTC().Format(time.RFC3339))
	},
}

func init() {
	time_cmd.Cmd.AddCommand(Cmd)
}
