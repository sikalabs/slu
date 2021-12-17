package list

import (
	"fmt"

	parent_cmd "github.com/sikalabs/slu/cmd/digitalocean/auth"
	"github.com/sikalabs/slu/config"

	"github.com/spf13/cobra"
)

var FlagShowTokens bool

var Cmd = &cobra.Command{
	Use:     "list",
	Short:   "List stored DigitalOcean accounts",
	Aliases: []string{"ls"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		s := config.ReadSecrets()
		for _, do := range s.DigitalOcean {
			if FlagShowTokens {
				fmt.Printf("%s: %s\n", do.Alias, do.Token)
			} else {
				fmt.Printf("%s: REDACTED\n", do.Alias)
			}
		}
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
	Cmd.PersistentFlags().BoolVarP(
		&FlagShowTokens,
		"show-tokens",
		"t",
		false,
		"Show DigitalOcean tokens",
	)
}
