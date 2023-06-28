package create_oidc_user

import (
	k8s_cmd "github.com/sikalabs/slu/cmd/k8s"
	"github.com/sikalabs/slu/utils/k8s_oidc_utils"

	"github.com/spf13/cobra"
)

var FlagName string
var FlagIssuerUrl string
var FlagClientId string
var FlagClientSecret string
var FlagDry bool

var Cmd = &cobra.Command{
	Use:   "create-oidc-user",
	Short: "Create OIDC User for kubelogin",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		k8s_oidc_utils.CreateOidcUser(
			FlagName, FlagIssuerUrl,
			FlagClientId, FlagClientSecret,
			FlagDry,
		)
	},
}

func init() {
	k8s_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVar(
		&FlagName,
		"name",
		"",
		"Name of user in kubeconfig",
	)
	Cmd.MarkFlagRequired("name")
	Cmd.Flags().StringVar(
		&FlagIssuerUrl,
		"issuer-url",
		"",
		"OIDC Issuer url (eg.: https://sso.sikalabs.com/realm/sikalabs)",
	)
	Cmd.MarkFlagRequired("issuer-url")
	Cmd.Flags().StringVar(
		&FlagClientId,
		"client-id",
		"",
		"OIDC client-id",
	)
	Cmd.MarkFlagRequired("client-id")
	Cmd.Flags().StringVar(
		&FlagClientSecret,
		"client-secret",
		"",
		"OIDC Client Secret",
	)
	Cmd.Flags().BoolVar(
		&FlagDry,
		"dry",
		false,
		"Dry run",
	)
}
