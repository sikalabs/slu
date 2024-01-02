package setup_runner_utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sikalabs/slu/utils/exec_utils"
)

func SetupGitlabRunnerDocker(gitlabUrl, token, hostname string, concurrency int, dryRun bool) error {
	var err error
	var args []string

	args = []string{
		"pull", "-q", "gitlab/gitlab-runner:latest",
	}
	if dryRun {
		printCommand("docker", args)
	} else {
		err = exec_utils.ExecOut("docker", args...)
		if err != nil {
			return err
		}
	}

	var etcVolume string
	var buildsVolume string
	etcVolume = "gitlab-runner-etc:/etc/gitlab-runner"
	buildsVolume = "gitlab-runner-builds:/builds"
	args = []string{
		"run", "-d",
		"--name", "gitlab-runner",
		"--restart", "always",
		"-v", etcVolume,
		"-v", buildsVolume,
		"-v", "/var/run/docker.sock:/var/run/docker.sock",
		"-v", "/etc/hosts:/etc/hosts",
		"gitlab/gitlab-runner:latest",
	}
	if dryRun {
		printCommand("docker", args)
	} else {
		err = exec_utils.ExecOut("docker", args...)
		if err != nil {
			return err
		}
	}

	args = []string{
		"exec", "gitlab-runner",
		"gitlab-runner", "register",
		"--non-interactive",
		"--url", gitlabUrl,
		"--token", token,
		"--name", hostname,
		"--executor", "docker",
		"--docker-pull-policy", "if-not-present",
		"--docker-image", "docker:git",
		"--docker-volumes", etcVolume,
		"--docker-volumes", buildsVolume,
		"--docker-volumes", "/var/run/docker.sock:/var/run/docker.sock",
	}
	if dryRun {
		printCommand("docker", args)
	} else {
		err = exec_utils.ExecOut("docker", args...)
		if err != nil {
			return err
		}
	}

	if concurrency != 1 {
		args = []string{
			"exec", "gitlab-runner",
			"sed", "-i", "s+concurrent = 1+concurrent = " + strconv.Itoa(concurrency) + "+g",
			"/etc/gitlab-runner/config.toml",
		}
		if dryRun {
			printCommand("docker", args)
		} else {
			err = exec_utils.ExecOut("docker", args...)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func printCommand(name string, args []string) {
	fmt.Print(name)
	for _, arg := range args {
		if strings.Contains(arg, " ") {
			arg = "\"" + arg + "\""
		}
		fmt.Print(" " + arg)
	}
	fmt.Println()
}
