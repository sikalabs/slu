package initial_password

import (
	"encoding/json"
	"fmt"

	argocd_cmd "github.com/sikalabs/slu/cmd/argocd"
	"github.com/sikalabs/slu/utils/argocd_utils"

	"github.com/spf13/cobra"
)

var FlagJson bool
var CmdFlagNamespace string

var Cmd = &cobra.Command{
	Use:     "initial-password",
	Short:   "Get ArgoCD Initial Passowrd",
	Aliases: []string{"ip"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		password := argocd_utils.ArgoCDGetInitialPassword(CmdFlagNamespace)

		if FlagJson {
			outJson, err := json.Marshal(password)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(outJson))
		} else {
			fmt.Println(password)
		}
	},
}

func init() {
	argocd_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&CmdFlagNamespace,
		"namespace",
		"n",
		"argocd",
		"ArgoCD Namespace",
	)
	Cmd.PersistentFlags().BoolVar(
		&FlagJson,
		"json",
		false,
		"Format output to JSON",
	)
}
