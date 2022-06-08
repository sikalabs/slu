package sh_utils

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
)

func HandleError(err error) {
	log.Fatalln(err)
}

func execOutDir(dir string, command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	return err
}

func ExecShOutDir(dir string, script string) error {
	return execOutDir(dir, "/bin/sh", "-c", script)
}

func ExecShOutHome(script string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	return ExecShOutDir(home, script)
}

func File(file_path, content string) {
	home, err := os.UserHomeDir()
	if err != nil {
		HandleError(err)
	}
	err = ioutil.WriteFile(path.Join(home, file_path), []byte(content), 0644)
	if err != nil {
		HandleError(err)
	}
}
