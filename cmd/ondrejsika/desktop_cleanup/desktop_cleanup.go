package desktop_cleanup

import (
	"fmt"
	"log"
	"time"

	parentcmd "github.com/sikalabs/slu/cmd/ondrejsika"
	"github.com/sikalabs/slu/utils/exec_utils"
	"github.com/spf13/cobra"
)

var FlagYesGoBuildCache bool
var FlagYesGoPkgModCache bool
var FlagYesYarnCache bool

var ListSh []string
var ListRm []string

var Cmd = &cobra.Command{
	Use:   "desktop-cleanup",
	Short: "Clean up desktop",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		// Prepare cleanup script
		registerSh("brew cleanup")
		registerRm(".minikube/cache")
		if FlagYesYarnCache {
			registerRm("./Library/Caches/Yarn/*")
		}
		registerRm("./Library/Caches/pip/*")
		if FlagYesGoBuildCache {
			registerRm("./Library/Caches/go-build/*")
		}
		registerRm("./Library/Caches/Homebrew/downloads/*")
		if FlagYesGoPkgModCache {
			registerRm("./go/pkg/mod/cache/*")
		}
		registerRm(".nvm/.cache/*")

		// Review cleanup script
		for _, script := range ListSh {
			fmt.Println(script)
		}
		for _, rmParam := range ListRm {
			fmt.Println("rm -rf", rmParam)
		}

		// Wait
		fmt.Println("Wait for 10 seconds... cancel using ctrl+c")
		time.Sleep(10 * time.Second)

		// Do cleanup
		for _, script := range ListSh {
			sh(script)
		}
		for _, rmParam := range ListRm {
			rm(rmParam)
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

func registerSh(s string) {
	ListSh = append(ListSh, s)
}

func registerRm(rmParam string) {
	ListRm = append(ListRm, rmParam)
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
