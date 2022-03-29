package add

import (
	parent_cmd "github.com/sikalabs/slu/cmd/digitalocean/auth"
	"github.com/sikalabs/slu/config"

	"github.com/spf13/cobra"
)

var FlagAlias string

var Cmd = &cobra.Command{
	Use:     "use-context",
	Short:   "Use DigitalOcean Account in context",
	Aliases: []string{"use"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		co := config.ReadState()
		co.DigitalOcean.CurrentContext = FlagAlias
		config.WriteState(co)
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
