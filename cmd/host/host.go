package host

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var FlagIPv4Only bool
var FlagIPv6Only bool
var FlagOneAddress bool

var Cmd = &cobra.Command{
	Use:   "host <domain>",
	Short: "Get IP from domain",
	Args:  cobra.ExactArgs(1),
	Run: func(c *cobra.Command, args []string) {
		var ipsStr []string
		ips, err := net.LookupIP(args[0])
		if err != nil {
			log.Fatalln(err)
		}
		for _, ip := range ips {
			ipStr := ip.String()
			if FlagIPv4Only {
				if strings.Contains(ipStr, ":") {
					continue
				}
			}
			if FlagIPv6Only {
				if strings.Contains(ipStr, ".") {
					continue
				}
			}
			ipsStr = append(ipsStr, ipStr)
			if FlagOneAddress {
				break
			}
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
	Cmd.Flags().BoolVarP(
		&FlagIPv4Only,
		"ipv4-only",
		"4",
		false,
		"Return IPv4 records only",
	)
	Cmd.Flags().BoolVarP(
		&FlagIPv6Only,
		"ipv6-only",
		"6",
		false,
		"Return IPv6 records only",
	)
	Cmd.Flags().BoolVar(
		&FlagOneAddress,
		"one",
		false,
		"Return only one IP address",
	)
}
