package create_makefile

import (
	"log"
	"os"

	parent_cmd "github.com/sikalabs/slu/cmd/scripts"
	"github.com/sikalabs/slu/utils/exec_utils"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "create-makefile",
	Short: "Create Makefile with optional include of Makefile.local.mk and commit it",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		if _, err := os.Stat("Makefile"); err == nil {
			log.Fatalln("Makefile already exists")
		}
		err := os.WriteFile("Makefile", []byte("-include Makefile.local.mk\n"), 0644)
		if err != nil {
			log.Fatalln(err)
		}
		exec_utils.ExecOut("git", "add", "Makefile")
		exec_utils.ExecOut(
			"git", "commit", "-n", "-m",
			"feat(Makefile): Create Makefile with optional include of Makefile.local.mk (gitignored)")
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}
