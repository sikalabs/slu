package example

import (
	"fmt"

	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var FlagName string

func init() {
	root.RootCmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(&FlagName, "name", "n", "World", "Name to greet")
}

var Cmd = &cobra.Command{
	Use:   "example",
	Short: "This is an example command, you can cerate your own command from this",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		example(FlagName)
	},
}

func example(name string) {
	fmt.Printf("Hello, %s!\n", name)
}
