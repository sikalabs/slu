package password

import (
	"fmt"

	random_cmd "github.com/sikalabs/slu/cmd/random"
	"github.com/sikalabs/slu/utils/random_utils"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "password",
	Short:   "Generate random password",
	Aliases: []string{"pwd", "passwd", "pass"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		fmt.Println(
			random_utils.RandomString(4, random_utils.LOWER) + "-" +
				random_utils.RandomString(4, random_utils.UPPER) + "-" +
				random_utils.RandomString(4, random_utils.DIGITS) + "-" +
				random_utils.RandomString(4, ""),
		)
	},
}

func init() {
	random_cmd.Cmd.AddCommand(Cmd)
}
