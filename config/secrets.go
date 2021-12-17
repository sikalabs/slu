package config

import (
	"os"
	"path"

	"github.com/sikalabs/slu/utils/json_utils"
)

func GetSluSecretFilePath() (string, error) {
	fileName := ".slu.secrets.json"

	// From environment variable SLU_CONFIG_DIR
	configDirEnv := os.Getenv("SLU_CONFIG_DIR")
	if configDirEnv != "" {
		return path.Join(configDirEnv, fileName), nil
	}

	// ~/.slu.secrets.json
	home, err := os.UserHomeDir()
	if err != nil {
		return "", nil
	}
	return path.Join(home, fileName), nil
}

func ReadSecrets() SluSecrets {
	configFile, _ := GetSluSecretFilePath()
	var s SluSecrets
	json_utils.ReadJsonFile(configFile, &s)
	return s
}

func WriteSecrets(s SluSecrets) {
	configFile, _ := GetSluSecretFilePath()
	json_utils.WriteJsonFile(configFile, s)
}
