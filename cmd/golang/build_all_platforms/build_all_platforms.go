package build_all_platforms

import (
	"fmt"
	"log"
	"strings"

	parent_cmd "github.com/sikalabs/slu/cmd/golang"
	"github.com/sikalabs/slu/utils/exec_utils"

	"github.com/spf13/cobra"
)

var FlagDryRun bool

var Cmd = &cobra.Command{
	Use:     "build-all-platforms",
	Short:   "Build for All Platforms",
	Aliases: []string{"build-all"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		binName := getBinNameOrDie()
		scripts := []string{
			"GOOS=linux GOARCH=amd64 go build -o " + binName + "_linux_amd64",
			"GOOS=linux GOARCH=arm64 go build -o " + binName + "_linux_arm64",
			"GOOS=darwin GOARCH=amd64 go build -o " + binName + "_darwin_amd64",
			"GOOS=darwin GOARCH=arm64 go build -o " + binName + "_darwin_arm64",
			"GOOS=windows GOARCH=amd64 go build -o " + binName + "_windows_amd64",
			"GOOS=windows GOARCH=arm64 go build -o " + binName + "_windows_arm64",
		}
		for _, script := range scripts {
			fmt.Println(script)
			if FlagDryRun {
				continue
			}
			err := exec_utils.ExecShOut(script)
			if err != nil {
				log.Fatalln(err)
			}
		}
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
	Cmd.PersistentFlags().BoolVar(
		&FlagDryRun,
		"dry-run",
		false,
		"Print commands, but do not execute them",
	)
}

func getBinNameOrDie() string {
	output, err := exec_utils.ExecStr("go", "list")
	if err != nil {
		log.Fatalln(err)
	}
	output = strings.ReplaceAll(output, "\n", "")
	split := strings.Split(output, "/")
	binName := split[len(split)-1]
	return binName
}
