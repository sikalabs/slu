package iceland

import (
	"fmt"
	"net/http"

	"github.com/ondrejsika/go-iceland"
	parentcmd "github.com/sikalabs/slu/cmd/web_server"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "iceland",
	Short: "Webserver with pictures from Iceland",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		server()
	},
}

func init() {
	parentcmd.Cmd.AddCommand(Cmd)
}

func server() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/jpeg")
		w.WriteHeader(http.StatusOK)
		w.Write(iceland.ICELAND_RIVER_AT_POOL_2022)
	})
	fmt.Println("Listen on 0.0.0.0:8000, see http://127.0.0.1:8000")
	http.ListenAndServe(":8000", nil)
}
