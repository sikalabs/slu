package host

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "host <domain>",
	Short: "Get IP from domain",
	Args:  cobra.ExactArgs(1),
	Run: func(c *cobra.Command, args []string) {
		var ipsStr []string
		ips, _ := net.LookupIP(args[0])
		for _, ip := range ips {
			ipStr := ip.String()
			ipsStr = append(ipsStr, ipStr)
		}

		if root.RootCmdFlagJson {
			outJson, err := json.Marshal(ipsStr)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(outJson))
		} else {
			for _, ip := range ipsStr {
				fmt.Println(ip)
			}
		}
	},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
