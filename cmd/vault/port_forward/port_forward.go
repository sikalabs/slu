package port_forward

import (
	"fmt"

	parent_cmd "github.com/sikalabs/slu/cmd/vault"
	"github.com/sikalabs/slu/utils/sh_utils"
	"github.com/spf13/cobra"
)

var FlagDry bool
var FlagNamespace string

var Cmd = &cobra.Command{
	Use:     "port-forward",
	Short:   "Vault Port Forward",
	Aliases: []string{"pf"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		sh(`echo See: http://127.0.0.1:8200`, FlagDry)
		sh(`kubectl port-forward \
--namespace `+FlagNamespace+` \
svc/vault 8200:8200`, FlagDry)
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagNamespace,
		"namespace",
		"n",
		"vault",
		"Kubernetes Namespace",
	)
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
