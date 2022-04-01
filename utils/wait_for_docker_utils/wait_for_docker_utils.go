package wait_for_docker_utils

import (
	"time"

	"github.com/sikalabs/slu/utils/docker_utils"
	"github.com/sikalabs/slu/utils/wait_for_utils"
)

func WaitForDocker(timeout int) {
	wait_for_utils.WaitFor(
		timeout, 100*time.Millisecond,
		func() (bool, bool, string, error) {
			_, err := docker_utils.Ping()
			if err == nil {
				return wait_for_utils.WaitForResponseSucceeded("Docker is running")
			}
			return wait_for_utils.WaitForResponseWaiting(err.Error())
		},
	)
}
