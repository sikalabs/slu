package example_server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	random_cmd "github.com/sikalabs/slu/cmd/random"
	"github.com/sikalabs/slu/utils/random_utils"
	"github.com/sikalabs/slu/version"
	"github.com/spf13/cobra"
)

var FlagPort int

var Cmd = &cobra.Command{
	Use:   "server",
	Short: "Server with random data endpoints",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		portStr := strconv.Itoa(FlagPort)
		hostname, _ := os.Hostname()
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "[slu "+version.Version+"] Server with random data endpoints! %s %s \n", hostname, portStr)
			fmt.Printf("RemoteAddr=%s\n", r.RemoteAddr)
		})

		http.HandleFunc("/v1/slu_random_password", func(w http.ResponseWriter, r *http.Request) {
			password, err := random_utils.RandomPassword()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			fmt.Fprintf(w, "%s\n", password)
		})

		fmt.Println("[slu " + version.Version + "] Server started on 0.0.0.0:" + portStr + ", see http://127.0.0.1:" + portStr)
		http.ListenAndServe(":"+portStr, nil)
	},
}

func init() {
	random_cmd.Cmd.AddCommand(Cmd)
	Cmd.PersistentFlags().IntVarP(
		&FlagPort,
		"port",
		"p",
		8000,
		"Listen on port",
	)
}
