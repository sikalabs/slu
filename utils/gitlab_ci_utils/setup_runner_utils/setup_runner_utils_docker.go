package setup_runner_utils

import (
	"strconv"

	"github.com/sikalabs/slu/utils/exec_utils"
)

func SetupGitlabRunnerDocker(gitlabUrl, registrationToken, hostname string, concurency int) error {
	var err error

	err = exec_utils.ExecOut(
		"docker", "pull", "-q", "gitlab/gitlab-runner:latest",
	)
	if err != nil {
		return err
	}

	var etcVolume string
	var buildsVolume string
	etcVolume = "gitlab-runner-etc:/etc/gitlab-runner"
	buildsVolume = "gitlab-runner-builds:/builds"
	err = exec_utils.ExecOut(
		"docker",
		"run", "-d",
		"--name", "gitlab-runner",
		"--restart", "always",
		"-v", etcVolume,
		"-v", buildsVolume,
		"-v", "/var/run/docker.sock:/var/run/docker.sock",
		"-v", "/etc/hosts:/etc/hosts",
		"gitlab/gitlab-runner:latest",
	)
	if err != nil {
		return err
	}

	err = exec_utils.ExecOut(
		"docker", "exec", "gitlab-runner",
		"gitlab-runner", "register",
		"--non-interactive",
		"--url", gitlabUrl,
		"--registration-token", registrationToken,
		"--name", hostname,
		"--executor", "docker",
		"--docker-pull-policy", "if-not-present",
		"--docker-image", "docker:git",
		"--docker-volumes", etcVolume,
		"--docker-volumes", buildsVolume,
		"--docker-volumes", "/var/run/docker.sock:/var/run/docker.sock",
	)
	if err != nil {
		return err
	}

	if concurency != 1 {
		err = exec_utils.ExecOut(
			"docker", "exec", "gitlab-runner",
			"sed", "-i", "s+concurrent = 1+concurrent = "+strconv.Itoa(concurency)+"+g",
			"/etc/gitlab-runner/config.toml",
		)
		if err != nil {
			return err
		}
	}

	return nil
}
