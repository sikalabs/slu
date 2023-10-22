package upload

import (
	parent_cmd "github.com/sikalabs/slu/cmd/google_drive"
	"github.com/sikalabs/slu/utils/google_drive_utils"

	"github.com/spf13/cobra"
)

var FlagClientId string
var FlagClientSecret string
var FlagAccessToken string

var Cmd = &cobra.Command{
	Use:     "upload",
	Short:   "Upload file to Goole Drive",
	Aliases: []string{"u"},
	Args:    cobra.ExactArgs(1),
	Run: func(c *cobra.Command, args []string) {
		google_drive_utils.Upload(
			FlagClientId,
			FlagClientSecret,
			FlagAccessToken,
			args[0],
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
	Cmd.Flags().StringVar(
		&FlagAccessToken,
		"access-token",
		"",
		"Google Drive Access Token",
	)
}
