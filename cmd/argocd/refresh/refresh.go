package refresh

import (
	argocd_cmd "github.com/sikalabs/slu/cmd/argocd"
	"github.com/sikalabs/slu/utils/argocd_utils"

	"github.com/spf13/cobra"
)

var FlagServerAddr string
var FlagPassword string
var FlagInsecure bool
var FlagCFAccessServiceTokenName string

var Cmd = &cobra.Command{
	Use:     "refresh <app>",
	Short:   "Refresh ArgoCD Application (get + refresh)",
	Aliases: []string{"r"},
	Args:    cobra.ExactArgs(1),
	Run: func(c *cobra.Command, args []string) {
		appName := args[0]

		// Lazy-load the ArgoCD domain if server flag is not provided
		serverAddr := FlagServerAddr
		if serverAddr == "" {
			argoCDDomain, err := argocd_utils.ArgoCDGetDomain(FlagCFAccessServiceTokenName)
			if err != nil {
				cobra.CheckErr(err)
			}
			serverAddr = argoCDDomain
		}

		password := FlagPassword
		if password == "" {
			password = argocd_utils.ArgoCDGetInitialPassword("argocd")
		}
		token := argocd_utils.ArgoCDGetToken(
			c.Context(),
			serverAddr,
			FlagInsecure,
			"admin",
			password,
			FlagCFAccessServiceTokenName,
		)
		argocd_utils.ArgoCDRefresh(
			c.Context(),
			serverAddr,
			FlagInsecure,
			token,
			appName,
			FlagCFAccessServiceTokenName,
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
	Cmd.Flags().StringVar(
		&FlagCFAccessServiceTokenName,
		"service-token-name",
		"",
		"Cloudflare Access Service Token Name (in Vault)",
	)
}
