package mon

import (
	"log"
	"os"
	"strconv"

	"github.com/sikalabs/slu/utils/mail_utils"
	"github.com/sikalabs/slu/utils/telegram_utils"
)

func Mon(config MonConfig) {
	hostname, err := os.Hostname()
	if err != nil {
		log.Default().Printf("Error getting hostname: %v", err)
		hostname = "unknown-host"
	}
	usage, err := getDiskUsage()
	if err != nil {
		log.Printf("Error getting disk usage: %v", err)
	}
	log.Printf(
		"Disk Usage: Total: %.2f GB, Used: %.2f GB, Free: %.2f GB, Used Percent: %.2f%%",
		usage.TotalGb, usage.UsedGb, usage.FreeGb, usage.UsedPercent,
	)
	if usage.UsedPercent > config.DiskUsageAlertThreshold {
		log.Printf(
			"!!! Disk usage alert: Used percent %.2f%% exceeds threshold %.2f%%, sending alerts",
			usage.UsedPercent,
			config.DiskUsageAlertThreshold,
		)
		for _, toEmail := range config.ToEmails {

			err := mail_utils.SendSimpleMail(
				config.SMTPServer,
				strconv.Itoa(config.SMTPPort),
				config.SMTPUsername,
				config.SMTPPassword,
				config.FromEmail,
				toEmail,
				"⚠️ [slu mon] "+hostname+": disk usage alert !!!",
				`Disk usage alert on `+hostname+`

Disk usage `+usage.UsedPercentStr+" exceeds threshold "+strconv.FormatFloat(config.DiskUsageAlertThreshold, 'f', 2, 64)+`%.

- Free: `+usage.FreeGbStr+`
- Used: `+usage.UsedGbStr+`
- Total: `+usage.TotalGbStr+`

-- slu mon
`,
			)
			if err != nil {
				log.Printf("Error sending email: %v", err)
			}
		}

		for _, chatID := range config.TelegramChatIDs {
			err := telegram_utils.TelegramSendMessage(
				config.TelegramBotToken,
				chatID,
				`⚠️ [slu mon] `+hostname+`: disk usage alert !!!

Disk usage `+usage.UsedPercentStr+" exceeds threshold "+strconv.FormatFloat(config.DiskUsageAlertThreshold, 'f', 2, 64)+`%.

- Free: `+usage.FreeGbStr+`
- Used: `+usage.UsedGbStr+`
- Total: `+usage.TotalGbStr+`

-- slu mon
`,
			)
			if err != nil {
				log.Printf("Error sending Telegram message: %v", err)
			}
		}
	}
}
