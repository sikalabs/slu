package install_hello_world

import (
	"fmt"

	parent_cmd "github.com/sikalabs/slu/cmd/scripts/kubernetes"
	"github.com/sikalabs/slu/utils/sh_utils"
	"github.com/spf13/cobra"
)

var FlagDry bool
var FlagHost string

var Cmd = &cobra.Command{
	Use:     "install-hello-world",
	Short:   "Install sikalabs/hello-world (hello-world-server)",
	Aliases: []string{"ihw"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		sh(`helm upgrade --install \
		hello-world hello-world \
	--repo https://helm.sikalabs.io \
	--create-namespace \
	--namespace hello-world \
	--set host=`+FlagHost+` \
	--wait`, FlagDry)
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
	Cmd.Flags().StringVar(
		&FlagHost,
		"host",
		"",
		"public hostname of the hello-world server",
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
