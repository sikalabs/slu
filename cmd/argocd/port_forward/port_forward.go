package port_forward

import (
	"fmt"

	parent_cmd "github.com/sikalabs/slu/cmd/argocd"
	"github.com/sikalabs/slu/utils/sh_utils"
	"github.com/spf13/cobra"
)

var FlagDry bool
var FlagNamespace string

var Cmd = &cobra.Command{
	Use:     "port-forward",
	Short:   "ArgoCD Port Forward",
	Aliases: []string{"pf"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		sh(`echo See: https://127.0.0.1:8443`, FlagDry)
		sh(`kubectl port-forward \
--namespace `+FlagNamespace+` \
svc/argocd-server 8443:443`, FlagDry)
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagNamespace,
		"namespace",
		"n",
		"argocd",
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
