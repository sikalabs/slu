package get

import (
	"log"
	"os"

	parent_cmd "github.com/sikalabs/slu/cmd/elasticsearch"
	"github.com/sikalabs/slu/utils/exec_utils"

	"github.com/spf13/cobra"
)

var FlagVerbose bool

var Cmd = &cobra.Command{
	Use:   "get <nodes/shards/...>",
	Short: "Get Nodes/Shards/... from ElasticSearch API",
	Args:  cobra.ExactArgs(1),
	Run: func(c *cobra.Command, args []string) {
		conn := os.Getenv("ELASTICSEARCH_CONNECTION")
		if conn == "" {
			log.Fatalf("ELASTICSEARCH_CONNECTION env var is not set")
		}
		url := conn + "/_cat/" + args[0]
		if FlagVerbose {
			url += "?v"
		}
		exec_utils.ExecHomeOut("curl", url)
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
	Cmd.PersistentFlags().BoolVarP(
		&FlagVerbose,
		"verbose",
		"v",
		false,
		"Verbose output",
	)
}
