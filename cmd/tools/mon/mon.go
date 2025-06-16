package mon

import (
	"log"
	"time"

	parent_cmd "github.com/sikalabs/slu/cmd/tools"
	"github.com/sikalabs/slu/pkg/tools/mon"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "mon",
	Short: "Simple VM monitoring tool",
	Run: func(c *cobra.Command, args []string) {
		for {
			mon.Mon()
			log.Println("Sleeping for 10 minutes")
			time.Sleep(10 * time.Minute)
		}
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}
