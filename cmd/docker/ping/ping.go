package ping

import (
	"fmt"
	"log"
	"os"

	parent_cmd "github.com/sikalabs/slu/cmd/docker"
	"github.com/sikalabs/slu/utils/docker_utils"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "ping",
	Short: "Check if Docker is running",
	Run: func(c *cobra.Command, args []string) {
		ok, err := docker_utils.Ping()
		if ok {
			fmt.Println("OK")
			os.Exit(0)
		}
		log.Fatalln(err)
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}
