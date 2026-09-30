package who_am_i

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/digitalocean/godo"
	parent_cmd "github.com/sikalabs/slu/cmd/digitalocean"
	"github.com/sikalabs/slu/utils/config_utils"
	"github.com/spf13/cobra"
)

var FlagAlias string

var Cmd = &cobra.Command{
	Use:     "who-am-i",
	Short:   "Show DigitalOcean account info for current token",
	Aliases: []string{"wai"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		token := getToken()
		client := godo.NewFromToken(token)
		account, _, err := client.Account.Get(context.TODO())
		if err != nil {
			log.Fatal(err)
		}
		data, err := json.MarshalIndent(map[string]*godo.Account{"account": account}, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(data))
	},
}

func getToken() string {
	if FlagAlias != "" {
		do := config_utils.GetDigitalOceanAccountByAlias(FlagAlias)
		if do == nil {
			log.Fatalf("No credentials found for alias %s", FlagAlias)
		}
		return do.Token
	}
	if token := os.Getenv("DIGITALOCEAN_TOKEN"); token != "" {
		return token
	}
	do := config_utils.GetCurrentDigitalOceanAccount()
	if do == nil {
		log.Fatal("No credentials or context found (use --alias, DIGITALOCEAN_TOKEN or slu do auth use-context)")
	}
	return do.Token
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagAlias,
		"alias",
		"a",
		"",
		"Use specific account (instead of DIGITALOCEAN_TOKEN or selected account in context)",
	)
}
