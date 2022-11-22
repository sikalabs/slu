package install_maildev

import (
	parent_cmd "github.com/sikalabs/slu/cmd/scripts/kubernetes"
	"github.com/sikalabs/slu/utils/k8s_scripts"
	"github.com/spf13/cobra"
)

var FlagDry bool
var FlagNamespace string
var FlagDomain string

var Cmd = &cobra.Command{
	Use:     "install-maildev",
	Short:   "Install Maildev",
	Aliases: []string{"imd"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		k8s_scripts.InstallMaildevDomain(FlagNamespace, FlagDomain, FlagDry)
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagNamespace,
		"namespace",
		"n",
		"maildev",
		"Kubernetes Namespace",
	)
	Cmd.Flags().BoolVar(
		&FlagDry,
		"dry",
		false,
		"Dry run",
	)
	Cmd.Flags().StringVarP(
		&FlagDomain,
		"domain",
		"d",
		"",
		"Domain of Maildev instance",
	)
	Cmd.MarkFlagRequired("domain")
}
