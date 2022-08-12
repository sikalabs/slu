package latest_release

import (
	"fmt"
	"log"
	"strings"

	parent_cmd "github.com/sikalabs/slu/cmd/github"
	"github.com/sikalabs/slu/utils/github_utils"

	"github.com/spf13/cobra"
)

var FlagNamespace string

var Cmd = &cobra.Command{
	Use:     "latest-release <user>/<repo>",
	Short:   "Get latest release of a repository",
	Aliases: []string{"lr"},
	Args:    cobra.ExactArgs(1),
	RunE: func(c *cobra.Command, args []string) error {
		argS := strings.Split(args[0], "/")
		if len(argS) != 2 {
			return fmt.Errorf("invalid argument: %s, required format is <user>/<repo>", args[0])
		}
		release, err := github_utils.GetLatestReleaseE(argS[0], argS[1])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(release)
		return nil
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}
