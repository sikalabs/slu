package get_token

import (
	parent_cmd "github.com/sikalabs/slu/cmd/google_drive"
	"github.com/sikalabs/slu/utils/google_drive_utils"

	"github.com/spf13/cobra"
)

var FlagClientId string
var FlagClientSecret string

var Cmd = &cobra.Command{
	Use:     "get-token",
	Short:   "Get Access Token to Goole Drive API",
	Aliases: []string{"u"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		google_drive_utils.GetToken(
			FlagClientId,
			FlagClientSecret,
		)
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVar(
		&FlagClientId,
		"client-id",
		"",
		"Google Drive Client ID",
	)
	Cmd.MarkFlagRequired("client-id")
	Cmd.Flags().StringVar(
		&FlagClientSecret,
		"client-secret",
		"",
		"Google Drive Client Secret",
	)
	Cmd.MarkFlagRequired("client-secret")
}
