package url

import (
	"fmt"
	"log"

	argocd_cmd "github.com/sikalabs/slu/cmd/argocd"
	"github.com/sikalabs/slu/utils/argocd_utils"

	"github.com/spf13/cobra"
)

var FlagNamespace string

var Cmd = &cobra.Command{
	Use:   "url",
	Short: "Get ArgoCD URL",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		domain, err := argocd_utils.ArgoCDGetDomain(FlagNamespace)
		if err != nil {
			log.Fatal(err)
		}

		url := "https://" + domain
		fmt.Println(url)
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
