package desktop_cleanup

import (
	"fmt"
	"log"
	"os"
	"time"

	parentcmd "github.com/sikalabs/slu/cmd/ondrejsika"
	"github.com/sikalabs/slu/utils/docker_utils"
	"github.com/sikalabs/slu/utils/exec_utils"
	"github.com/spf13/cobra"
)

var FlagDryRun bool
var FlagYesGoBuildCache bool
var FlagYesGoPkgModCache bool
var FlagYesYarnCache bool
var FlagTerraformPluginDir bool
var FlagBrewCache bool
var FlagNoDockerPrune bool

var ListSh []string
var ListRm []string

var Cmd = &cobra.Command{
	Use:   "desktop-cleanup",
	Short: "Clean up desktop",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		// Prepare cleanup script
		registerRm(".minikube/cache")
		if FlagYesYarnCache {
			registerRm("./Library/Caches/Yarn/*")
		}
		registerRm("./Library/Caches/pip/*")
		if FlagYesGoBuildCache {
			registerRm("./Library/Caches/go-build/*")
		}
		if FlagBrewCache {
			registerRm("./Library/Caches/Homebrew/*")
		}
		if FlagYesGoPkgModCache {
			registerRm("./go/pkg/mod/cache/*")
		}
		registerRm(".nvm/.cache/*")
		if FlagTerraformPluginDir {
			registerRm(".terraform-plugin-cache/*")
		}
		registerRm("./Library/Caches/lima/")
		if !FlagNoDockerPrune {
			dockerUp, _ := docker_utils.Ping()
			if dockerUp {
				registerSh("docker system prune --force")
			}
		}

		// Review cleanup script
		for _, rmParam := range ListRm {
			fmt.Println("rm -rf", rmParam)
		}
		for _, script := range ListSh {
			fmt.Println(script)
		}

		if FlagDryRun {
			os.Exit(0)
		}

		// Wait
		fmt.Println("Wait for 10 seconds... cancel using ctrl+c")
		time.Sleep(10 * time.Second)

		// Do cleanup
		for _, rmParam := range ListRm {
			rm(rmParam)
		}
		for _, script := range ListSh {
			sh(script)
		}
	},
}

func init() {
	parentcmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().BoolVar(
		&FlagDryRun,
		"dry-run",
		false,
		"Dry run",
	)
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
	Cmd.Flags().BoolVar(
		&FlagTerraformPluginDir,
		"terraform-plugin-dir",
		false,
		"Remove Terraform providers Dir (rm -rf  ~/terraform-plugin-cache/*)",
	)
	Cmd.Flags().BoolVar(
		&FlagBrewCache,
		"brew-cache",
		false,
		"Cleanup Brew cache (rm -rf ~/Library/Caches/Homebrew/*)",
	)
	Cmd.Flags().BoolVar(
		&FlagNoDockerPrune,
		"no-docker-prune",
		false,
		"Do not run docker system prune",
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
