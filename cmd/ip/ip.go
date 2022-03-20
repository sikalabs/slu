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

var FlagJson bool

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

		if FlagJson {
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
	Cmd.PersistentFlags().BoolVar(
		&FlagJson,
		"json",
		false,
		"Format output to JSON",
	)
}
