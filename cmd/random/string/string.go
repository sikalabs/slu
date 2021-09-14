package initial_password

import (
	"fmt"

	random_cmd "github.com/sikalabs/slu/cmd/random"
	"github.com/sikalabs/slu/utils/random_utils"

	"github.com/spf13/cobra"
)

var FlagLenght int

var Cmd = &cobra.Command{
	Use:     "string",
	Short:   "Get Random string",
	Aliases: []string{"str", "s"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		fmt.Println(random_utils.RandomString(FlagLenght, ""))
	},
}

func init() {
	random_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().IntVarP(
		&FlagLenght,
		"length",
		"l",
		16,
		"Length of random string",
	)
}
