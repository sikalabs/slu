package wait_for_tcp_utils

import (
	"net"
	"time"

	"github.com/sikalabs/slu/utils/wait_for_utils"
)

func WaitForTcp(timeout int, addr string) {
	wait_for_utils.WaitFor(
		timeout, 100*time.Millisecond,
		func() (bool, bool, string, error) {
			_, err := net.DialTimeout("tcp", addr, 100*time.Millisecond)
			if err == nil {
				return wait_for_utils.WaitForResponseSucceeded("Connected")
			}
			return wait_for_utils.WaitForResponseWaiting(err.Error())
		},
	)

}
