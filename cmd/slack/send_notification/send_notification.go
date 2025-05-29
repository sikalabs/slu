package send_notification

import (
	"log"

	parent_cmd "github.com/sikalabs/slu/cmd/slack"
	"github.com/sikalabs/slu/utils/slack_utils"

	"github.com/spf13/cobra"
)

var FlagUrl string
var FlagSource string
var FlagTitle string
var FlagMessage string
var FlagStatus string

var Cmd = &cobra.Command{
	Use:     "send-notification",
	Short:   "Send notification to Slack channel",
	Aliases: []string{"sn"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		err := slack_utils.SendNotification(slack_utils.Notification{
			URLs:    []string{FlagUrl},
			Source:  FlagSource,
			Title:   FlagTitle,
			Message: FlagMessage,
			Status:  FlagStatus,
		})

		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagUrl,
		"url",
		"u",
		"",
		"Slack webhook Url",
	)
	Cmd.MarkFlagRequired("url")

	Cmd.Flags().StringVarP(
		&FlagTitle,
		"title",
		"t",
		"",
		"Title of the notification",
	)
	Cmd.MarkFlagRequired("title")

	Cmd.Flags().StringVarP(
		&FlagSource,
		"source",
		"s",
		"",
		"Source of the notification event (e.g. 'MySQL PROD', 'Vault DEV', etc.)",
	)
	Cmd.MarkFlagRequired("source")

	Cmd.Flags().StringVarP(
		&FlagMessage,
		"message",
		"m",
		"",
		"Message of the notification (e.g. 'Database connection lost', 'New secret created', etc.)",
	)
	Cmd.MarkFlagRequired("message")

	Cmd.Flags().StringVarP(
		&FlagStatus,
		"status",
		"S",
		"INFO",
		"Status of the notification, only INFO, OK, WARN or ERR are allowed",
	)
}
