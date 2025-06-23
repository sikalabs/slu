package password_reset

import (
	"fmt"

	parent_cmd "github.com/sikalabs/slu/cmd/keycloak"
	"github.com/sikalabs/slu/pkg/utils/keycloak_utils"

	"github.com/spf13/cobra"
)

var FlagKeycloakUrl string
var FlagAdminUsername string
var FlagAdminPassword string
var FlagRealm string
var FlagUsername string
var FlagNewPassword string

var Cmd = &cobra.Command{
	Use:     "password-reset",
	Short:   "Reset Keycloak User Password",
	Aliases: []string{"pr"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		fmt.Println(
			FlagKeycloakUrl,
			FlagAdminUsername,
			FlagAdminPassword,
			FlagRealm,
			FlagUsername,
			FlagNewPassword,
		)
		keycloak_utils.PasswordResetOrDie(
			FlagKeycloakUrl,
			FlagAdminUsername,
			FlagAdminPassword,
			FlagRealm,
			FlagUsername,
			FlagNewPassword,
		)
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVar(
		&FlagKeycloakUrl,
		"keycloak-url",
		"",
		"Keycloak URL",
	)
	Cmd.MarkFlagRequired("keycloak-url")
	Cmd.Flags().StringVar(
		&FlagAdminUsername,
		"admin-username",
		"",
		"Keycloak Admin User",
	)
	Cmd.MarkFlagRequired("admin-username")
	Cmd.Flags().StringVar(
		&FlagAdminPassword,
		"admin-password",
		"",
		"Keycloak Admin Password",
	)
	Cmd.MarkFlagRequired("admin-password")
	Cmd.Flags().StringVar(
		&FlagRealm,
		"realm",
		"",
		"Keycloak Realm",
	)
	Cmd.MarkFlagRequired("realm")
	Cmd.Flags().StringVar(
		&FlagUsername,
		"username",
		"",
		"Keycloak Username",
	)
	Cmd.MarkFlagRequired("username")
	Cmd.Flags().StringVar(
		&FlagNewPassword,
		"new-password",
		"",
		"New Keycloak Password",
	)
	Cmd.MarkFlagRequired("new-password")
}
