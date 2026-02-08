package send_file

import (
	"log"
	"os"

	parentcmd "github.com/sikalabs/slu/cmd/telegram"
	"github.com/sikalabs/slu/utils/telegram_utils"
	"github.com/spf13/cobra"
)

var FlagBotToken string
var FlagChatID int64
var FlagMessage string
var FlagFile string

var Cmd = &cobra.Command{
	Use:   "send-file",
	Short: "Send a file to a Telegram chat",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		err := telegram_utils.TelegramSendFile(FlagBotToken, FlagChatID, FlagFile, FlagMessage)
		if err != nil {
			log.Fatalln(err)
		}
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
		&FlagFile,
		"file",
		"f",
		"",
		"Path to file to send",
	)
	Cmd.MarkFlagRequired("file")
	Cmd.Flags().StringVarP(
		&FlagMessage,
		"message",
		"m",
		"",
		"Message to send with the file (optional)",
	)
}
