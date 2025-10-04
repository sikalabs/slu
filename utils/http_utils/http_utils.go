package http_utils

import (
	"fmt"
	"io"
	"net/http"

	"github.com/sikalabs/slu/internal/error_utils"
)

func UrlGetToStringE(url string) (string, error) {
	var client http.Client
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		return string(bodyBytes), nil
	}
	return "", fmt.Errorf("status code is not 200 OK: %d", resp.StatusCode)
}

func UrlGetToString(url string) string {
	body, err := UrlGetToStringE(url)
	error_utils.HandleError(err, "Failed to get URL content")
	return body
}
