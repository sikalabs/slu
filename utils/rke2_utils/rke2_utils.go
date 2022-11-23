package rke2_utils

import (
	_ "embed"

	"log"

	"github.com/sikalabs/slu/utils/template_utils"
)

//go:embed rke2-config-master.yml.tpl
var MasterYamlTemplate string

type TemplateInventory struct {
	IsFirstMaster bool
	ServerDomain  string
	Token         string
	TlsSans       []string
}

func GenerateMasterConfig(
	domain string,
	token string,
	sans []string,
	isFirstMaster bool,
) string {
	out, err := template_utils.TemplateFromStringToString(
		"rke2-config-master.yml",
		MasterYamlTemplate,
		TemplateInventory{
			ServerDomain:  domain,
			Token:         token,
			TlsSans:       sans,
			IsFirstMaster: isFirstMaster,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	return out
}

func GenerateWorkerConfig(
	token string,
	sans []string,
) string {
	return ""
}
