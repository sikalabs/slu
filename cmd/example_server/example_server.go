package example_server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/sikalabs/slu/cmd/root"
	"github.com/sikalabs/slu/version"
	"github.com/spf13/cobra"
)

var FlagPort int

var Cmd = &cobra.Command{
	Use:   "example-server",
	Short: "Run example web server",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		portStr := strconv.Itoa(FlagPort)
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			hostname, _ := os.Hostname()
			fmt.Fprintf(w, "[slu "+version.Version+"] Example HTTP Server! %s %s \n", hostname, portStr)
		})
		fmt.Println("[slu " + version.Version + "] Server started on 0.0.0.0:" + portStr + ", see http://127.0.0.1:" + portStr)
		http.ListenAndServe(":"+portStr, nil)
	},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
	Cmd.PersistentFlags().IntVarP(
		&FlagPort,
		"port",
		"p",
		8000,
		"Listen on port",
	)
}
