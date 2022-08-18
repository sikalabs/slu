package domain

import (
	"fmt"

	argocd_cmd "github.com/sikalabs/slu/cmd/argocd"
	"github.com/sikalabs/slu/utils/argocd_utils"

	"github.com/spf13/cobra"
)

var FlagNamespace string

var Cmd = &cobra.Command{
	Use:     "domain",
	Short:   "Get ArgoCD Domain",
	Aliases: []string{"dom", "d"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		fmt.Println(argocd_utils.ArgoCDGetDomainOrDie(FlagNamespace))
	},
}

func init() {
	argocd_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagNamespace,
		"namespace",
		"n",
		"argocd",
		"ArgoCD Namespace",
	)
}
