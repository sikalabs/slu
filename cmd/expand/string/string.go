package string

import (
	"encoding/json"
	"fmt"
	"os"

	expandcmd "github.com/sikalabs/slu/cmd/expand"

	"github.com/spf13/cobra"
)

var FlagJson bool
var CmdFlagSource string

var Cmd = &cobra.Command{
	Use:   "string",
	Short: "Expand environment variables in string",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		if FlagJson {
			outJson, err := json.Marshal(os.ExpandEnv(CmdFlagSource))
			if err != nil {
				panic(err)
			}
			fmt.Println(string(outJson))
		} else {
			fmt.Println(os.ExpandEnv(CmdFlagSource))
		}
	},
}

func init() {
	expandcmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&CmdFlagSource,
		"source",
		"s",
		"",
		"Source template string",
	)
	Cmd.MarkFlagRequired("source")
	Cmd.PersistentFlags().BoolVar(
		&FlagJson,
		"json",
		false,
		"Format output to JSON",
	)
}
