package human_readable_password

import (
	"fmt"

	random_cmd "github.com/sikalabs/slu/cmd/random"
	"github.com/sikalabs/slu/utils/random_utils"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "human-readable-password",
	Short:   "Generate random human-readable password using BIP39 mnemonic",
	Aliases: []string{"hrp"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		password := random_utils.RandomHumanReadablePassword()
		fmt.Println(password)
	},
}

func init() {
	random_cmd.Cmd.AddCommand(Cmd)
}
