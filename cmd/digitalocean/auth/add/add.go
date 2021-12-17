package add

import (
	"fmt"

	parent_cmd "github.com/sikalabs/slu/cmd/digitalocean/auth"
	"github.com/sikalabs/slu/config"

	"github.com/spf13/cobra"
)

var FlagAlias string
var FlagToken string

var Cmd = &cobra.Command{
	Use:   "add",
	Short: "Add DigitalOcean Account",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		s := config.ReadSecrets()

		var updated bool = false
		for i, do := range s.DigitalOcean {
			fmt.Println(do)
			if do.Alias == FlagAlias {
				s.DigitalOcean[i].Token = FlagToken
				updated = true
			}
		}

		if !updated {
			s.DigitalOcean = append(s.DigitalOcean, config.SluSecretsDigitalOcean{
				Alias: FlagAlias,
				Token: FlagToken,
			})
		}

		config.WriteSecrets(s)
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
	Cmd.PersistentFlags().StringVarP(
		&FlagAlias,
		"alias",
		"a",
		"",
		"Alias of DigitalOcean Account",
	)
	Cmd.MarkPersistentFlagRequired("alias")
	Cmd.PersistentFlags().StringVarP(
		&FlagToken,
		"token",
		"t",
		"",
		"DigitalOcean Token",
	)
	Cmd.MarkPersistentFlagRequired("token")
}
