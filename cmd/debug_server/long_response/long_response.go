package long_response

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	parent_cmd "github.com/sikalabs/slu/cmd/debug_server"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "long-response",
	Short:   "Server with long response time",
	Aliases: []string{"long", "lr", "lore"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			hostname, _ := os.Hostname()
			sleepTime, err := strconv.Atoi(strings.ReplaceAll(r.URL.Path, "/", ""))
			if err != nil {
				fmt.Println(err)
				fmt.Println("set sleep time to 0s (default)")
			}
			time.Sleep(time.Duration(sleepTime) * time.Second)
			fmt.Fprintf(w, "[slu-debug-server] Response after %ds! %s \n", sleepTime, hostname)
		})
		fmt.Println("Server started.")
		http.ListenAndServe(":8000", nil)
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}
