package telegram_utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func TelegramGetLastChatID(botToken string) (int64, error) {
	type Chat struct {
		ID int64 `json:"id"`
	}

	type Message struct {
		Chat Chat `json:"chat"`
	}

	type Update struct {
		Message Message `json:"message"`
	}

	type Response struct {
		Ok     bool     `json:"ok"`
		Result []Update `json:"result"`
	}

	resp, err := http.Get("https://api.telegram.org/bot" + botToken + "/getUpdates")
	if err != nil {
		return 0, fmt.Errorf("getting updates failed: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("reading response body failed: %v", err)
	}

	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		return 0, fmt.Errorf("unmarshalling response failed: %v", err)
	}

	if len(response.Result) == 0 {
		return 0, fmt.Errorf("no updates found")
	}

	lastChatID := response.Result[len(response.Result)-1].Message.Chat.ID
	return lastChatID, nil
}
