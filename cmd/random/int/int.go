package int

import (
	"fmt"

	random_cmd "github.com/sikalabs/slu/cmd/random"
	"github.com/sikalabs/slu/utils/random_utils"

	"github.com/spf13/cobra"
)

var FlagMin int
var FlagMax int

var Cmd = &cobra.Command{
	Use:     "int",
	Short:   "Get Random int from interval [min,max)",
	Aliases: []string{"int", "i"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		fmt.Println(random_utils.RandomInt(FlagMin, FlagMax))
	},
}

func init() {
	random_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().IntVar(
		&FlagMin,
		"min",
		0,
		"Minimum of random interval",
	)
	Cmd.Flags().IntVar(
		&FlagMax,
		"max",
		100,
		"Maximum of random interval",
	)
}
