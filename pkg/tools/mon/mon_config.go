package mon

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type MonConfig struct {
	SMTPServer   string   `json:"smtp_server" yaml:"smtp_server"`
	SMTPPort     int      `json:"smtp_port" yaml:"smtp_port"`
	SMTPUsername string   `json:"smtp_username" yaml:"smtp_username"`
	SMTPPassword string   `json:"smtp_password" yaml:"smtp_password"`
	FromEmail    string   `json:"from_email" yaml:"from_email"`
	ToEmails     []string `json:"to_emails" yaml:"to_emails"`

	DiskUsageAlertThreshold float64 `json:"disk_usage_alert_threshold" yaml:"disk_usage_alert_threshold"`
}

func ReadMonConfig(configFilePath string) (MonConfig, error) {
	var config MonConfig

	file, err := os.ReadFile(configFilePath)
	if err != nil {
		log.Printf("Error reading config file: %v", err)
		return config, err
	}

	if err := yaml.Unmarshal(file, &config); err != nil {
		log.Printf("Error unmarshalling config file: %v", err)
		return config, err
	}

	return config, nil
}
