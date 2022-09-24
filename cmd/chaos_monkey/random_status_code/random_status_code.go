package random_status_code

import (
	"os"

	parent_cmd "github.com/sikalabs/slu/cmd/chaos_monkey"
	"github.com/sikalabs/slu/utils/random_utils"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "random-status-code",
	Aliases: []string{"rsc"},
	Short:   "Return random status code",
	Run: func(c *cobra.Command, args []string) {
		ok := random_utils.RandomBool()
		if ok {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}
