package example_server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "example-server",
	Short: "Run example web server",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			hostname, _ := os.Hostname()
			fmt.Fprintf(w, "[slu] Example HTTP Server! %s \n", hostname)
		})
		fmt.Println("Server started.")
		http.ListenAndServe(":8000", nil)
	},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
