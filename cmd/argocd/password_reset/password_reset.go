package password_reset

import (
	"log"

	argocd_cmd "github.com/sikalabs/slu/cmd/argocd"
	"github.com/sikalabs/slu/utils/exec_utils"

	"github.com/spf13/cobra"
)

var FlagNamespace string

var Cmd = &cobra.Command{
	Use:     "password-reset",
	Short:   "Reset ArgoCD Admin Passowrd",
	Aliases: []string{"pr"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		// TODO: Rewrite using Go
		var err error
		err = exec_utils.ExecOut("sh", "-c", `kubectl patch secret argocd-secret  -p '{"data": {"admin.password": null, "admin.passwordMtime": null}}' -n `+FlagNamespace)
		if err != nil {
			log.Fatalln(err)
		}
		err = exec_utils.ExecOut("sh", "-c", `kubectl delete pod -l app.kubernetes.io/component=server -n `+FlagNamespace)
		if err != nil {
			log.Fatalln(err)
		}
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
