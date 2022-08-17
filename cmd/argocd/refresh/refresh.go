package refresh

import (
	argocd_cmd "github.com/sikalabs/slu/cmd/argocd"
	"github.com/sikalabs/slu/utils/argocd_utils"

	"github.com/spf13/cobra"
)

var FlagServerAddr string
var FlagPassword string
var FlagInsecure bool

var Cmd = &cobra.Command{
	Use:     "refresh <app>",
	Short:   "Refresh ArgoCD Application (get + refresh)",
	Aliases: []string{"r"},
	Args:    cobra.ExactArgs(1),
	Run: func(c *cobra.Command, args []string) {
		appName := args[0]
		password := FlagPassword
		if password == "" {
			password = argocd_utils.ArgoCDGetInitialPassword("argocd")
		}
		token := argocd_utils.ArgoCDGetToken(
			c.Context(),
			FlagServerAddr,
			FlagInsecure,
			"admin",
			password,
		)
		argocd_utils.ArgoCDRefresh(
			c.Context(),
			FlagServerAddr,
			FlagInsecure,
			token,
			appName,
		)
	},
}

func init() {
	argocd_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagServerAddr,
		"server",
		"s",
		"",
		"ArgoCD server address (host:port)",
	)
	Cmd.MarkFlagRequired("server")
	Cmd.Flags().BoolVar(
		&FlagInsecure,
		"insecure",
		false,
		"Insecure connection",
	)
	Cmd.Flags().StringVarP(
		&FlagPassword,
		"password",
		"p",
		"",
		"ArgoCD server password",
	)
}
