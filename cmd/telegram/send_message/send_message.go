package send_message

import (
	"os"

	parentcmd "github.com/sikalabs/slu/cmd/telegram"
	"github.com/sikalabs/slu/utils/telegram_utils"
	"github.com/spf13/cobra"
)

var FlagBotToken string
var FlagChatID int64
var FlagMessage string

var Cmd = &cobra.Command{
	Use:   "send-message",
	Short: "Send a message to a Telegram chat",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		telegram_utils.TelegramSendMessageMarkdown(FlagBotToken, FlagChatID, FlagMessage)
	},
}

func init() {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	parentcmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagBotToken,
		"bot-token",
		"t",
		botToken,
		"Telegram Bot token, can be set via TELEGRAM_BOT_TOKEN env var",
	)
	if botToken == "" {
		Cmd.MarkFlagRequired("bot-token")
	}
	Cmd.Flags().Int64VarP(
		&FlagChatID,
		"chat-id",
		"c",
		0,
		"Chat ID",
	)
	Cmd.MarkFlagRequired("chat-id")
	Cmd.Flags().StringVarP(
		&FlagMessage,
		"message",
		"m",
		"",
		"Message to send",
	)
	Cmd.MarkFlagRequired("message")
}
