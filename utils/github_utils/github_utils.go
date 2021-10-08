package github_utils

import (
	"encoding/json"
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
	var err error

	resp, err := http.Get("https://api.github.com/repos/" + user + "/" + repo + "/releases/latest")
	handleError(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	handleError(err)
	data := LatestReleaseResponse{}

	err = json.Unmarshal(body, &data)
	handleError(err)

	return data.TagName
}
