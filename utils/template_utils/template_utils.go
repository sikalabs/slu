package template_utils

import (
	"bytes"
	"html/template"
)

func TemplateFromStringToString(
	templateName string,
	templateBody string,
	data any,
) (string, error) {
	var err error
	var out bytes.Buffer

	tmpl, err := template.New(templateName).Parse(templateBody)
	if err != nil {
		return "", err
	}

	if err := tmpl.Execute(&out, data); err != nil {
		return "", err
	}

	return out.String(), nil
}
