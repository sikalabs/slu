package desktop_cleanup

import (
	"log"

	parentcmd "github.com/sikalabs/slu/cmd/ondrejsika"
	"github.com/sikalabs/slu/utils/exec_utils"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "desktop-cleanup",
	Short: "Clean up desktop",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		rm(".minikube/cache")
		rm("./Library/Caches/Yarn/*")
		rm("./Library/Caches/pip/*")
		rm("./Library/Caches/go-build/*")
		rm("./Library/Caches/Homebrew/downloads/*")
	},
}

func init() {
	parentcmd.Cmd.AddCommand(Cmd)
}

func sh(script string) {
	err := exec_utils.ExecShHomeOut(script)
	if err != nil {
		log.Fatalln(err)
	}
}

func rm(path string) {
	sh("rm -rf " + path)
}
