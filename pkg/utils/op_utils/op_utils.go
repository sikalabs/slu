package op_utils

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"

	"github.com/sikalabs/slu/utils/exec_utils"
)

type OPItem struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Vault struct {
		ID string `json:"id"`
	} `json:"vault"`
}

func CheckOpBinaryExists() error {
	_, err := exec.LookPath("op")
	if err != nil {
		return fmt.Errorf("1Password CLI (op) not found in PATH")
	}
	return nil
}

func CheckOpBinaryExistsOrDie() {
	err := CheckOpBinaryExists()
	if err != nil {
		log.Fatal(err)
	}
}

func GetFileFrom1PasswordOrDie(vault, name, outputPath string) {
	err := GetFileFrom1Password(vault, name, outputPath)
	if err != nil {
		log.Fatalf("Failed to get file from 1Password: %v", err)
	}
}

func GetFileFrom1Password(vault, name, outputPath string) error {
	var err error

	log.Printf("Get %s from 1Password vault %s", name, vault)

	// Get document from 1Password
	err = sh(`op document get "` + name + `" --vault "` + vault + `" --out-file "` + outputPath + `"`)
	if err != nil {
		return err
	}

	log.Println()
	log.Printf("Successfully saved to %s", outputPath)
	log.Println()

	return nil
}

func SaveFileTo1PasswordOrDie(vault, filePath string) {
	err := SaveFileTo1Password(vault, filePath)
	if err != nil {
		log.Fatalf("Failed to save file to 1Password: %v", err)
	}
}

func SaveFileTo1Password(vault, filePath string) error {
	var err error
	fileName := filepath.Base(filePath)

	log.Printf("Save %s to 1Password vault %s", fileName, vault)

	// Create document in 1Password
	err = sh(`op document create --vault "` + vault + `" --title "` + fileName + `" "` + filePath + `"`)
	if err != nil {
		return err
	}

	log.Println()
	log.Printf("Successfully saved %s to 1Password vault %s", fileName, vault)
	log.Println()

	return nil
}

func Get(vault, name, field string) (string, error) {
	log.Printf("Get field %s from %s in vault %s", field, name, vault)

	// Use op read with secret reference format
	reference := "op://" + vault + "/" + name + "/" + field
	output, err := shCapture(`op read "` + reference + `" --no-newline`)
	if err != nil {
		return "", err
	}

	return output, nil
}

func sh(cmd string) error {
	log.Printf("RUN: %s", cmd)
	return exec_utils.ExecShOut(cmd)
}

func shCapture(cmd string) (string, error) {
	log.Printf("RUN: %s", cmd)
	cmd2 := exec.Command("sh", "-c", cmd)
	out, err := cmd2.CombinedOutput()
	return string(out), err
}
