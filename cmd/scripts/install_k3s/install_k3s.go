package install_k3s

import (
	parent_cmd "github.com/sikalabs/slu/cmd/scripts"
	"github.com/sikalabs/slu/utils/k8s_scripts"
	"github.com/spf13/cobra"
)

var FlagDry bool

var Cmd = &cobra.Command{
	Use:     "install-k3s",
	Short:   "Install k3s without Traefik",
	Aliases: []string{"ik3s"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		k8s_scripts.InstallK3s(FlagDry)
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
