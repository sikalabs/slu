package up

import (
	"fmt"
	"regexp"

	"github.com/sikalabs/slu/cmd/root"
	"github.com/sikalabs/slu/utils/exec_utils"
	"github.com/sikalabs/slu/utils/file_utils"
	"github.com/spf13/cobra"
)

var FlagJson bool

var Cmd = &cobra.Command{
	Use:   "up [flags] [-- <args>, ...]",
	Short: "Try to run make up, docker-compose up, go run main.go",
	Run: func(c *cobra.Command, args []string) {
		if checkMakeUp() {
			fmt.Println("slu up: make up")
			exec_utils.ExecOut("make", append([]string{"up"}, args...)...)
		} else if checkDockerComposeUp() {
			fmt.Println("slu up: docker compose up")
			exec_utils.ExecOut("docker", append([]string{"compose", "up"}, args...)...)
		} else if checkGoRunMailGO() {
			fmt.Println("slu up: go run main.go")
			exec_utils.ExecOut("go", append([]string{"run", "main.go"}, args...)...)
		}
	},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}

func checkUpTargetInMakefile(makefile string) bool {
	r := regexp.MustCompile(`\nup:`)
	match := r.MatchString(makefile)
	return match
}

func checkMakeUp() bool {
	if !file_utils.FileExists("Makefile") {
		return false
	}
	makefile, err := file_utils.ReadFileToString("Makefile")
	if err != nil {
		return false
	}
	return checkUpTargetInMakefile(makefile)
}

func checkDockerComposeUp() bool {
	return file_utils.FileExists("docker-compose.yml") || file_utils.FileExists("docker-compose.yaml")
}

func checkGoRunMailGO() bool {
	return file_utils.FileExists("main.go")
}
