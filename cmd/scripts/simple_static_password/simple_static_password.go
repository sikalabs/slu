package simple_static_password

import (
	"fmt"

	parent_cmd "github.com/sikalabs/slu/cmd/scripts"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "simple-static-password",
	Aliases: []string{"ssp"},
	Short:   "Print simple static password `asdf_ASDF_1234`",
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		fmt.Println("asdf_ASDF_1234")
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}
