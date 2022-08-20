package set

import (
	"fmt"
	"log"
	"syscall"

	parent_cmd "github.com/sikalabs/slu/cmd/cloudflare/access/service_token"
	"github.com/sikalabs/slu/lib/vault_cfa_service_token"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

var FlagClientID string
var FlagClientSecret string

var Cmd = &cobra.Command{
	Use:   "set",
	Short: "Set Cloudflare Access Service Token",
	Args:  cobra.ExactArgs(1),
	Run: func(c *cobra.Command, args []string) {
		key := args[0]
		clientID := FlagClientID
		clientSecret := FlagClientSecret
		if clientID == "" {
			fmt.Print("CLIENT_ID:")
			clientID = readPasswordOrDie()
		}
		if clientSecret == "" {
			fmt.Print("CLIENT_SECRET:")
			clientSecret = readPasswordOrDie()
		}
		err := vault_cfa_service_token.Set(key, clientID, clientSecret)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}

func readPasswordOrDie() string {
	passwordBytes, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print("\n")
	return string(passwordBytes)
}
