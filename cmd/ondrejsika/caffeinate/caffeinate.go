package caffeinate

import (
	"log"

	parentcmd "github.com/sikalabs/slu/cmd/ondrejsika"
	"github.com/sikalabs/slu/utils/exec_utils"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "caffeinate",
	Short: "Prevent the Mac from sleeping (runs caffeinate -di)",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		err := exec_utils.ExecInOut("caffeinate", "-di")
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	parentcmd.Cmd.AddCommand(Cmd)
}
