package example_server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

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
		hostname, _ := os.Hostname()
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "[slu "+version.Version+"] Example HTTP Server! %s %s \n", hostname, portStr)
		})
		http.HandleFunc("/slow1s", func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(1 * time.Second)
			fmt.Fprintf(w, "[slu "+version.Version+"] Example HTTP Server (after 1s)! %s %s \n", hostname, portStr)
		})
		http.HandleFunc("/slow10s", func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(10 * time.Second)
			fmt.Fprintf(w, "[slu "+version.Version+"] Example HTTP Server (after 10s)! %s %s \n", hostname, portStr)
		})
		http.HandleFunc("/slow30s", func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(30 * time.Second)
			fmt.Fprintf(w, "[slu "+version.Version+"] Example HTTP Server (after 30s)! %s %s \n", hostname, portStr)
		})
		http.HandleFunc("/slow60s", func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(60 * time.Second)
			fmt.Fprintf(w, "[slu "+version.Version+"] Example HTTP Server (after 60s)! %s %s \n", hostname, portStr)
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
