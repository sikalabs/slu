package rm

import (
	parent_cmd "github.com/sikalabs/slu/cmd/digitalocean/auth"
	"github.com/sikalabs/slu/config"

	"github.com/spf13/cobra"
)

var FlagAlias string

var Cmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove DigitalOcean Account",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		s := config.ReadSecrets()
		dos := []config.SluSecretsDigitalOcean{}
		for _, do := range s.DigitalOcean {
			if do.Alias != FlagAlias {
				dos = append(dos, do)
			}
		}
		s.DigitalOcean = dos
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
}
