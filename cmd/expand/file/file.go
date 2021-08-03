package File

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	expandcmd "github.com/sikalabs/slut/cmd/expand"
	rootcmd "github.com/sikalabs/slut/cmd/root"

	"github.com/spf13/cobra"
)

var CmdFlagSource string

var Cmd = &cobra.Command{
	Use:   "file",
	Short: "Expand environment variables in file",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		source, err := ioutil.ReadFile(CmdFlagSource)
		if err != nil {
			panic(err)
		}

		if rootcmd.RootCmdFlagJson {
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
}
