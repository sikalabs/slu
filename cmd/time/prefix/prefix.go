package unix

import (
	"bufio"
	"fmt"
	"os"
	"time"

	time_cmd "github.com/sikalabs/slu/cmd/time"
	"github.com/sikalabs/slu/utils/time_utils"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "prefix",
	Short: "Prefix stdin with timestap",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		start := time.Now()
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			since := time.Since(start)
			fmt.Println(time_utils.DurationToString(since), s.Text())
		}
	},
}

func init() {
	time_cmd.Cmd.AddCommand(Cmd)
}
