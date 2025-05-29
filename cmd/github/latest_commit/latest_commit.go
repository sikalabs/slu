package latest_commit

import (
	"fmt"
	"log"
	"strings"

	parent_cmd "github.com/sikalabs/slu/cmd/github"
	"github.com/sikalabs/slu/utils/github_utils"

	"github.com/spf13/cobra"
)

var FlagShort bool

var Cmd = &cobra.Command{
	Use:     "latest-commit <user>/<repo>",
	Short:   "Get latest commit of a repository",
	Aliases: []string{"lc"},
	Args:    cobra.ExactArgs(1),
	RunE: func(c *cobra.Command, args []string) error {
		argS := strings.Split(args[0], "/")
		if len(argS) != 2 {
			return fmt.Errorf("invalid argument: %s, required format is <user>/<repo>", args[0])
		}
		commit, err := github_utils.GetLatestCommitE(argS[0], argS[1])
		if err != nil {
			log.Fatal(err)
		}

		if FlagShort {
			commit = commit[:7] // Shorten to first 7 characters
		}

		fmt.Println(commit)
		return nil
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().BoolVarP(&FlagShort, "short", "s", false, "Short commit hash")
}
