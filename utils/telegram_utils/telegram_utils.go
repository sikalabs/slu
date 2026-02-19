package telegram_utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func TelegramSendMessage(botToken string, chatID int64, message string) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)
	body, err := json.Marshal(map[string]string{
		"chat_id": fmt.Sprintf("%d", chatID),
		"text":    message,
	})
	if err != nil {
		return err
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return checkTelegramResponse(resp)
}

func TelegramSendMessageMarkdown(botToken string, chatID int64, message string) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)
	body, err := json.Marshal(map[string]string{
		"chat_id":    fmt.Sprintf("%d", chatID),
		"text":       message,
		"parse_mode": "MarkdownV2",
	})
	if err != nil {
		return err
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return checkTelegramResponse(resp)
}

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
		Ok          bool     `json:"ok"`
		Description string   `json:"description"`
		Result      []Update `json:"result"`
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

	if !response.Ok {
		return 0, fmt.Errorf("telegram API error: %s", response.Description)
	}

	if len(response.Result) == 0 {
		return 0, fmt.Errorf("no updates found")
	}

	lastChatID := response.Result[len(response.Result)-1].Message.Chat.ID
	return lastChatID, nil
}

func TelegramSendFile(botToken string, chatID int64, filePath string, message string, asPhoto bool) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// Add chat_id field
	_ = writer.WriteField("chat_id", fmt.Sprintf("%d", chatID))

	// Add caption/message
	_ = writer.WriteField("caption", message)

	// Add the file
	fileField := "document"
	if asPhoto {
		fileField = "photo"
	}
	part, err := writer.CreateFormFile(fileField, filepath.Base(filePath))
	if err != nil {
		return fmt.Errorf("failed to create form file: %v", err)
	}

	_, err = bytes.NewBuffer(nil).ReadFrom(file)
	if _, err = file.Seek(0, 0); err != nil {
		return fmt.Errorf("failed to reset file: %v", err)
	}

	_, err = file.WriteTo(part)
	if err != nil {
		return fmt.Errorf("failed to write file to part: %v", err)
	}

	writer.Close()

	apiMethod := "sendDocument"
	if asPhoto {
		apiMethod = "sendPhoto"
	}
	url := fmt.Sprintf("https://api.telegram.org/bot%s/%s", botToken, apiMethod)

	req, err := http.NewRequest("POST", url, &requestBody)
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %v", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	return checkTelegramResponse(resp)
}

func checkTelegramResponse(resp *http.Response) error {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("reading response body failed: %v", err)
	}
	var result struct {
		Ok          bool   `json:"ok"`
		Description string `json:"description"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return fmt.Errorf("unmarshalling response failed: %v", err)
	}
	if !result.Ok {
		return fmt.Errorf("telegram API error: %s", result.Description)
	}

	return nil
}
