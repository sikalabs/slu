package update_docker_images

import (
	"fmt"

	parent_cmd "github.com/sikalabs/slu/cmd/scripts/gitlab_ci"
	"github.com/sikalabs/slu/utils/sh_utils"
	"github.com/spf13/cobra"
)

var FlagDry bool

var Cmd = &cobra.Command{
	Use:     "update-docker-images",
	Short:   "Update Docker Images on Gitlab CI Runner",
	Aliases: []string{"udi"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		sh(`
echo Updating docker pull -q sikalabs/ci ...
docker pull -q sikalabs/ci
echo Updating docker pull -q sikalabs/ci-node ...
docker pull -q sikalabs/ci-node
`, FlagDry)
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().BoolVar(
		&FlagDry,
		"dry",
		false,
		"Dry run",
	)
}

func sh(script string, dry bool) {
	if dry {
		fmt.Println(script)
		return
	}
	err := sh_utils.ExecShOutDir("", script)
	if err != nil {
		sh_utils.HandleError(err)
	}
}
