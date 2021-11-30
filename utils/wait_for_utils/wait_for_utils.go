package wait_for_utils

import (
	"fmt"
	"log"
	"os"
	"time"
)

func WaitFor(
	timeout int,
	sleep time.Duration,
	waitFunc func() (bool, bool, string, error),
) {
	started := time.Now()
	for {
		succeeded, failed, message, err := waitFunc()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(message)

		if succeeded {
			os.Exit(0)
		}

		if failed {
			os.Exit(1)
		}

		if time.Since(started) > time.Duration(timeout*int(time.Second)) {
			os.Exit(1)
		}

		time.Sleep(sleep)
	}
}

func WaitForResponseSucceeded(message string) (bool, bool, string, error) {
	return true, false, message, nil
}

func WaitForResponseFailed(message string) (bool, bool, string, error) {
	return false, true, message, nil
}

func WaitForResponseWaiting(message string) (bool, bool, string, error) {
	return false, false, message, nil
}

func WaitForResponseError(err error) (bool, bool, string, error) {
	return false, false, err.Error(), err
}
