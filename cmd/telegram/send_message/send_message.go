package send_message

import (
	"os"
	"strconv"

	parentcmd "github.com/sikalabs/slu/cmd/telegram"
	"github.com/sikalabs/slu/pkg/utils/error_utils"
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
		err := telegram_utils.TelegramSendMessageMarkdown(FlagBotToken, FlagChatID, FlagMessage)
		error_utils.HandleError(err)
	},
}

func init() {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	var chatID int64
	if chatIDStr := os.Getenv("TELEGRAM_CHAT_ID"); chatIDStr != "" {
		chatID, _ = strconv.ParseInt(chatIDStr, 10, 64)
	}
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
		chatID,
		"Chat ID, can be set via TELEGRAM_CHAT_ID env var",
	)
	if chatID == 0 {
		Cmd.MarkFlagRequired("chat-id")
	}
	Cmd.Flags().StringVarP(
		&FlagMessage,
		"message",
		"m",
		"",
		"Message to send",
	)
	Cmd.MarkFlagRequired("message")
}
