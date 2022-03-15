package ip

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "ip",
	Short: "Get my current IP address (using checkip.amazonaws.com)",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		resp, err := http.Get("https://checkip.amazonaws.com/")
		if err != nil {
			panic(err)
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		ip := strings.TrimSuffix(string(body), "\n")

		if root.RootCmdFlagJson {
			outJson, err := json.Marshal(map[string]string{
				"ip": ip,
			})
			if err != nil {
				panic(err)
			}
			fmt.Println(string(outJson))
		} else {
			fmt.Printf("%s\n", ip)
		}
	},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
