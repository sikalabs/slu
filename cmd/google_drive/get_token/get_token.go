package get_token

import (
	"log"

	parent_cmd "github.com/sikalabs/slu/cmd/google_drive"
	"github.com/sikalabs/slu/utils/google_drive_utils"
	"github.com/sikalabs/slu/utils/vault_google_drive_utils"

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
		token := google_drive_utils.GetToken(FlagClientId, FlagClientSecret)
		err := vault_google_drive_utils.SetGoogleDriveUploadTokenSecrets(token, "secret/data/slu/google-drive-upload/token")
		if err != nil {
			log.Fatalln(err)
		}
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
	Cmd.Flags().StringVar(
		&FlagClientSecret,
		"client-secret",
		"",
		"Google Drive Client Secret",
	)
}
