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
	Use:   "ip-local",
	Short: "Get local IP from network device",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
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
	Cmd.MarkFlagRequired("interface")
}
