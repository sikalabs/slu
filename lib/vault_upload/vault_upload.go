package vault_upload

import (
	"github.com/sikalabs/slu/utils/vault_s3_utils"
)

func GetUploadSecrets() (string, string, string, string, string, error) {
	return vault_s3_utils.GetS3Secrets("secret/data/slu/upload")
}
