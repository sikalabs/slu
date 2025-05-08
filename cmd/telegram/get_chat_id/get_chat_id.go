package get_chat_id

import (
	"fmt"
	"log"

	parentcmd "github.com/sikalabs/slu/cmd/telegram"
	"github.com/sikalabs/slu/utils/telegram_utils"
	"github.com/spf13/cobra"
)

var FlagBotToken string
var FlagSendToChat bool

var Cmd = &cobra.Command{
	Use:   "get-chat-id",
	Short: "Get Telegram chat ID",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		chatID, err := telegram_utils.TelegramGetLastChatID(FlagBotToken)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(chatID)
		if FlagSendToChat {
			telegram_utils.TelegramSendMessage(FlagBotToken, chatID, fmt.Sprintf("`%d`", chatID))
		}
	},
}

func init() {
	parentcmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagBotToken,
		"bot-token",
		"t",
		FlagBotToken,
		"Telegram Bot token",
	)
	Cmd.MarkFlagRequired("bot-token")
	Cmd.Flags().BoolVarP(
		&FlagSendToChat,
		"send-to-chat",
		"s",
		FlagSendToChat,
		"Send message to chat",
	)
}
