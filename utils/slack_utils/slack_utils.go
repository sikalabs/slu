package slack_utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const (
	ColorGreen  = "#36a64f"
	ColorRed    = "#ff0000"
	ColorYellow = "#f5cc45"
	ColorBlue   = "#add8e6"
	SlIconURL   = "https://avatars.githubusercontent.com/u/16311959?s=200&v=4"
)

type StatusStyle struct {
	Color string
	Icon  string
}

type Notification struct {
	URLs    []string
	Source  string
	Status  string
	Title   string
	Message string
}

var StatusInfo = StatusStyle{
	Color: ColorBlue,
	Icon:  "ℹ️",
}

var StatusSuccess = StatusStyle{
	Color: ColorGreen,
	Icon:  "✅",
}

var StatusWarning = StatusStyle{
	Color: ColorYellow,
	Icon:  "⚠️",
}

var StatusError = StatusStyle{
	Color: ColorRed,
	Icon:  "❌",
}

func SendNotification(notification Notification) error {
	if len(notification.URLs) == 0 {
		return fmt.Errorf("no Slack webhook URLs provided")
	}

	statusStyles := map[string]StatusStyle{
		"INFO": StatusInfo,
		"OK":   StatusSuccess,
		"WARN": StatusWarning,
		"ERR":  StatusError,
	}

	statusStyle, exists := statusStyles[notification.Status]
	if !exists {
		return fmt.Errorf("unknown status: %s", notification.Status)
	}

	message := strings.ReplaceAll(notification.Message, `\n`, "\n")

	text := fmt.Sprintf("%s *%s %s*\n%s",
		statusStyle.Icon,
		notification.Title,
		statusStyle.Icon,
		message,
	)

	attachment := map[string]interface{}{
		"color":     statusStyle.Color,
		"text":      text,
		"mrkdwn_in": []string{"text"},
		"fallback":  notification.Title,
	}

	payload := map[string]interface{}{
		"attachments": []map[string]interface{}{attachment},
		"username":    notification.Source,
		"icon_url":    SlIconURL,
	}

	for _, url := range notification.URLs {
		body, err := json.Marshal(payload)
		if err != nil {
			return fmt.Errorf("failed to marshal JSON: %w", err)
		}

		resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))

		if err != nil {
			return fmt.Errorf("failed to send Slack notification: %w", err)
		}

		defer resp.Body.Close()
	}

	return nil
}
