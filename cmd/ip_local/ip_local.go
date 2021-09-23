package ip_local

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/sikalabs/slu/cmd/root"
	"github.com/sikalabs/slu/utils/ip_utils"
	"github.com/spf13/cobra"
)

var FlagInterfaceName string

var Cmd = &cobra.Command{
	Use:     "ip-local",
	Short:   "Get local IP from network device",
	Aliases: []string{"ipl"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		if FlagInterfaceName == "" {
			ips, err := ip_utils.GetIPFromInterfaces()
			if err != nil {
				log.Fatal(err)
			}
			if root.RootCmdFlagJson {
				outJson, err := json.Marshal(map[string]map[string]string{
					"ips": ips,
				})
				if err != nil {
					panic(err)
				}
				fmt.Println(string(outJson))
			} else {
				for interfaceName, ip := range ips {
					fmt.Printf("%s=%s\n", interfaceName, ip)
				}
			}
		} else {
			ip, err := ip_utils.GetIPFromInterface(FlagInterfaceName)
			if err != nil {
				log.Fatal(err)
			}

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
		}
	},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagInterfaceName,
		"interface",
		"i",
		"",
		"Interface name",
	)
}
