package install_cnpg

import (
	parent_cmd "github.com/sikalabs/slu/cmd/scripts/kubernetes"
	"github.com/sikalabs/slu/utils/k8s_scripts"
	"github.com/spf13/cobra"
)

var FlagDry bool
var FlagInstallOnly bool

var Cmd = &cobra.Command{
	Use:     "install-cnpg",
	Short:   "Install CNPG (Cloud Native Postgres)",
	Aliases: []string{"icnpg"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		k8s_scripts.InstallCNPG(FlagDry, FlagInstallOnly)
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
	Cmd.Flags().BoolVar(
		&FlagInstallOnly,
		"install-only",
		false,
		"Use helm install instead of helm upgrade --install",
	)
}
