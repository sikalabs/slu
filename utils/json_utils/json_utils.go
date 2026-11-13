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
	return WriteJsonFileWithMode(filename, jsonData, 0644)
}

// WriteJsonFileWithMode writes JSON to filename and makes sure the file has
// the given mode, also when the file already exists (os.WriteFile applies
// the mode only to newly created files).
func WriteJsonFileWithMode(filename string, jsonData interface{}, mode os.FileMode) error {
	var err error
	fileData, err := json.Marshal(jsonData)
	if err != nil {
		return err
	}
	err = os.Chmod(filename, mode)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	err = os.WriteFile(filename, fileData, mode)
	if err != nil {
		return err
	}
	return nil
}
