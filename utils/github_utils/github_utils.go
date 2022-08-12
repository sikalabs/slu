package github_utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type LatestReleaseResponse struct {
	TagName string `json:"tag_name"`
}

func GetLatestRelease(user, repo string) string {
	release, err := GetLatestReleaseE(user, repo)
	handleError(err)
	return release
}

func GetLatestReleaseE(user, repo string) (string, error) {
	var err error

	resp, err := http.Get("https://api.github.com/repos/" + user + "/" + repo + "/releases/latest")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("repository %s/%s does not exist", user, repo)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	data := LatestReleaseResponse{}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", err
	}

	return data.TagName, nil
}
