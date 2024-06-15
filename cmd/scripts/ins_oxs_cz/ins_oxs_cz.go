package ins_oxs_cz

import (
	"fmt"

	parent_cmd "github.com/sikalabs/slu/cmd/scripts"
	"github.com/sikalabs/slu/utils/sh_utils"
	"github.com/spf13/cobra"
)

var FlagDry bool

var Cmd = &cobra.Command{
	Use:     "ins-oxs-cz <script>",
	Aliases: []string{"ins"},
	Short:   "Run script from ins.oxs.cz",
	Args:    cobra.ExactArgs(1),
	Run: func(c *cobra.Command, args []string) {
		name := args[0]
		sh("curl -fsSL https://ins.oxs.cz/"+name+".sh | sh", FlagDry)
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().BoolVar(
		&FlagDry,
		"dry-run",
		false,
		"Dry run",
	)
}

func sh(script string, dry bool) {
	if dry {
		fmt.Println(script)
		return
	}
	err := sh_utils.ExecShOutDir("", script)
	if err != nil {
		sh_utils.HandleError(err)
	}
}
