package install_hello_world

import (
	parent_cmd "github.com/sikalabs/slu/cmd/scripts/kubernetes"
	"github.com/sikalabs/slu/utils/k8s_scripts"
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
		k8s_scripts.InstallHelloWorld(FlagHost, FlagDry)
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
