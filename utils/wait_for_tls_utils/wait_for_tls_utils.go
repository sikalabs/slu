package wait_for_tls_utils

import (
	"crypto/tls"
	"time"

	"github.com/sikalabs/slu/utils/wait_for_utils"
)

func WaitForTls(timeout int, addr string) {
	wait_for_utils.WaitFor(
		timeout, 100*time.Millisecond,
		func() (bool, bool, string, error) {
			_, err := tls.Dial("tcp", addr, &tls.Config{
				InsecureSkipVerify: false,
			})
			if err == nil {
				return wait_for_utils.WaitForResponseSucceeded("TLS certificate validated")
			}
			return wait_for_utils.WaitForResponseWaiting(err.Error())
		},
	)
}
