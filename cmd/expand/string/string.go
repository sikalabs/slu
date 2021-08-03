package string

import (
	"encoding/json"
	"fmt"
	"os"

	expandcmd "github.com/sikalabs/slut/cmd/expand"
	rootcmd "github.com/sikalabs/slut/cmd/root"

	"github.com/spf13/cobra"
)

var CmdFlagSource string

var Cmd = &cobra.Command{
	Use:   "string",
	Short: "Expand environment variables in string",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		if rootcmd.RootCmdFlagJson {
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
}
