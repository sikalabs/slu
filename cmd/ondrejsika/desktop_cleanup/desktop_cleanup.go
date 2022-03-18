package desktop_cleanup

import (
	"log"

	parentcmd "github.com/sikalabs/slu/cmd/ondrejsika"
	"github.com/sikalabs/slu/utils/exec_utils"
	"github.com/spf13/cobra"
)

var FlagYesGoBuildCache bool
var FlagYesGoPkgModCache bool
var FlagYesYarnCache bool

var Cmd = &cobra.Command{
	Use:   "desktop-cleanup",
	Short: "Clean up desktop",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		sh("brew cleanup")
		rm(".minikube/cache")
		if FlagYesYarnCache {
			rm("./Library/Caches/Yarn/*")
		}
		rm("./Library/Caches/pip/*")
		if FlagYesGoBuildCache {
			rm("./Library/Caches/go-build/*")
		}
		rm("./Library/Caches/Homebrew/downloads/*")
		if FlagYesGoPkgModCache {
			rm("./go/pkg/mod/cache/*")
		}
	},
}

func init() {
	parentcmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().BoolVar(
		&FlagYesGoBuildCache,
		"go-build-cache",
		false,
		"Cleanup GO build cache (rm -rf  ~/Library/Caches/go-build/*)",
	)
	Cmd.Flags().BoolVar(
		&FlagYesGoPkgModCache,
		"go-pkg-mod-cache",
		false,
		"Cleanup GO build cache (rm -rf  ~/go/pkg/mod/cache/*)",
	)
	Cmd.Flags().BoolVar(
		&FlagYesYarnCache,
		"yarn-cache",
		false,
		"Cleanup Yarn cache (rm -rf  ~/Library/Caches/Yarn/*)",
	)
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
