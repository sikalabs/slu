package systemd_utils

import (
	"github.com/sikalabs/slu/utils/template_utils"
)

const SYSTEMD_SERVICE_TEMPLATE = `[Unit]
Description={{.Description}}
ConditionPathExists={{.WorkingDirectory}}
After=network.target

[Service]
Type=simple
User={{.User}}
Group={{.Group}}
WorkingDirectory={{.WorkingDirectory}}

ExecStart={{.ExecStart}}
Restart=on-failure
RestartSec=10

[Install]
WantedBy=multi-user.target
`

func CreateSystemdServiceString(
	name string,
	description string,
	user string,
	group string,
	workingDirectory string,
	execStart string,
) (string, error) {
	out, err := template_utils.TemplateFromStringToString(
		"systemd_service",
		SYSTEMD_SERVICE_TEMPLATE,
		map[string]string{
			"Name":             name,
			"Description":      description,
			"User":             user,
			"Group":            group,
			"WorkingDirectory": workingDirectory,
			"ExecStart":        execStart,
		},
	)
	if err != nil {
		return "", err
	}
	return out, nil
}
