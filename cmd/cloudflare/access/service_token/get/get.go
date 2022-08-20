package get

import (
	"fmt"
	"log"

	parent_cmd "github.com/sikalabs/slu/cmd/cloudflare/access/service_token"
	"github.com/sikalabs/slu/lib/vault_cfa_service_token"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "get",
	Short: "Get Cloudflare Access Service Token",
	Args:  cobra.ExactArgs(1),
	Run: func(c *cobra.Command, args []string) {
		key := args[0]
		clientId, clientSecret, err := vault_cfa_service_token.Get(key)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("CLIENT_ID:     ", clientId)
		fmt.Println("CLIENT_SECRER: ", clientSecret)
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}
