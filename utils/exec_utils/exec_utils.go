package exec_utils

import (
	"os"
	"os/exec"
)

func ExecOut(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	return err
}

func ExecShOut(script string) error {
	return ExecOut("sh", "-c", script)
}

func ExecNoOut(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	err := cmd.Run()
	return err
}

func ExecHomeOut(command string, args ...string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	cmd := exec.Command(command, args...)
	cmd.Dir = home
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	return err
}

func ExecShHomeOut(script string) error {
	return ExecHomeOut("sh", "-c", script)
}

func ExecStr(command string, args ...string) (string, error) {
	out, err := exec.Command(command, args...).Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
