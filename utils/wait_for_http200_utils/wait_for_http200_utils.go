package wait_for_http200_utils

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sikalabs/slu/utils/wait_for_utils"
)

func WaitForHttp200(timeout int, url string) {
	client := &http.Client{
		Timeout: 1 * time.Second,
	}

	wait_for_utils.WaitFor(
		timeout, 100*time.Millisecond,
		func() (bool, bool, string, error) {
			resp, err := client.Get(url)
			if err != nil {
				return wait_for_utils.WaitForResponseWaiting(err.Error())
			}
			defer resp.Body.Close()

			if resp.StatusCode == http.StatusOK {
				return wait_for_utils.WaitForResponseSucceeded(fmt.Sprintf("HTTP %d", resp.StatusCode))
			}
			return wait_for_utils.WaitForResponseWaiting(fmt.Sprintf("HTTP %d", resp.StatusCode))
		},
	)
}
