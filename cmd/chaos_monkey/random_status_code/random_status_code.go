package random_status_code

import (
	"fmt"
	"os"

	parent_cmd "github.com/sikalabs/slu/cmd/chaos_monkey"
	"github.com/sikalabs/slu/utils/random_utils"
	"github.com/spf13/cobra"
)

var FlagVerbose bool

var Cmd = &cobra.Command{
	Use:     "random-status-code",
	Aliases: []string{"rsc"},
	Short:   "Return random status code",
	Run: func(c *cobra.Command, args []string) {
		ok := random_utils.RandomBool()
		if ok {
			if FlagVerbose {
				fmt.Println("exit code 0")
			}
			os.Exit(0)
		} else {
			fmt.Println("exit code 1")
			os.Exit(1)
		}
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
	Cmd.PersistentFlags().BoolVarP(
		&FlagVerbose,
		"verbose",
		"v",
		false,
		"Verbose output",
	)
}
