package config

import (
	"os"
	"path"

	"github.com/sikalabs/slu/utils/json_utils"
)

func GetSluConfigFilePath() (string, error) {
	fileName := ".slu.config.json"

	// From environment variable SLU_CONFIG_DIR
	configDirEnv := os.Getenv("SLU_CONFIG_DIR")
	if configDirEnv != "" {
		return path.Join(configDirEnv, fileName), nil
	}

	// ~/.slu.config.json
	home, err := os.UserHomeDir()
	if err != nil {
		return "", nil
	}
	return path.Join(home, fileName), nil
}

func ReadConfig() SluConfig {
	configFile, _ := GetSluConfigFilePath()
	var c SluConfig
	json_utils.ReadJsonFile(configFile, &c)
	return c
}

func WriteConfig(c SluConfig) {
	configFile, _ := GetSluConfigFilePath()
	json_utils.WriteJsonFile(configFile, c)
}
