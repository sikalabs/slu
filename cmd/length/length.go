package length

import (
	"encoding/json"
	"fmt"

	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var FlagJson bool

var Cmd = &cobra.Command{
	Use:     "length <string>",
	Short:   "Length of a string",
	Aliases: []string{"len"},
	Args:    cobra.ExactArgs(1),
	Run: func(c *cobra.Command, args []string) {
		l := len(args[0])
		if FlagJson {
			outJson, err := json.Marshal(l)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(outJson))
		} else {
			fmt.Println(l)
		}
	},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
	Cmd.PersistentFlags().BoolVar(
		&FlagJson,
		"json",
		false,
		"Format output to JSON",
	)
}
