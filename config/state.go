package config

import (
	"os"
	"path"

	"github.com/sikalabs/slu/utils/json_utils"
)

func GetSluStateFilePath() (string, error) {
	fileName := ".slu.state.json"

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

func ReadConfig() SluState {
	configFile, _ := GetSluStateFilePath()
	var c SluState
	json_utils.ReadJsonFile(configFile, &c)
	return c
}

func WriteConfig(c SluState) {
	configFile, _ := GetSluStateFilePath()
	json_utils.WriteJsonFile(configFile, c)
}
