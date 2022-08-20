package request_debug

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	parent_cmd "github.com/sikalabs/slu/cmd/debug_server"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "request-debug",
	Short:   "Return request degbug information",
	Aliases: []string{"req", "request", "rd"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			var err error
			rr := &Request{}
			rr.Method = r.Method
			rr.Headers = r.Header
			rr.URL = r.URL.String()
			rr.Body, err = ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			rrb, err := json.Marshal(rr)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(rrb)
			fmt.Println(string(rrb))
		})
		fmt.Println("Server started on 0.0.0.0:8000, see http://127.0.0.1:8000")
		http.ListenAndServe(":8000", nil)
	},
}

type Request struct {
	URL     string      `json:"url"`
	Method  string      `json:"method"`
	Headers http.Header `json:"headers"`
	Body    []byte      `json:"body"`
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}
