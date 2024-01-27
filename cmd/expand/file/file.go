package File

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
	Use:   "file",
	Short: "Expand environment variables in file",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		source, err := os.ReadFile(CmdFlagSource)
		if err != nil {
			panic(err)
		}

		if FlagJson {
			outJson, err := json.Marshal(os.ExpandEnv(string(source)))
			if err != nil {
				panic(err)
			}
			fmt.Println(string(outJson))
		} else {
			fmt.Println(os.ExpandEnv(string(source)))
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
		"Source template file",
	)
	Cmd.MarkFlagRequired("source")
	Cmd.PersistentFlags().BoolVar(
		&FlagJson,
		"json",
		false,
		"Format output to JSON",
	)
}
