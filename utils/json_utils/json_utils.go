package json_utils

import (
	"encoding/json"
	"os"
)

func ReadJsonFile(filename string, jsonData interface{}) error {
	var err error
	fileData, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(fileData, &jsonData)
	if err != nil {
		return err
	}
	return nil
}

func WriteJsonFile(filename string, jsonData interface{}) error {
	var err error
	fileData, err := json.Marshal(jsonData)
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, fileData, 0644)
	if err != nil {
		return err
	}
	return nil
}
