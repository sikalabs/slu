package github_utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/sikalabs/slu/internal/error_utils"
)

type LatestReleaseResponse struct {
	TagName string `json:"tag_name"`
}

type TagsResponse struct {
	Name string `json:"name"`
}

type CommitResponse struct {
	Sha string `json:"sha"`
}

func GetLatestRelease(user, repo string) string {
	release, err := GetLatestReleaseE(user, repo)
	error_utils.HandleError(err, "Failed to get latest release")
	return release
}

func GetLatestTag(user, repo string) string {
	release, err := GetLatestTagE(user, repo)
	error_utils.HandleError(err, "Failed to get latest tag")
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

	body, err := io.ReadAll(resp.Body)
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

func GetLatestTagE(user, repo string) (string, error) {
	var err error

	resp, err := http.Get("https://api.github.com/repos/" + user + "/" + repo + "/tags")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("repository %s/%s does not exist", user, repo)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	data := []TagsResponse{}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", err
	}

	if len(data) == 0 {
		return "", fmt.Errorf("repository %s/%s has no tags", user, repo)
	}

	return data[0].Name, nil
}

func GetLatestCommitE(user, repo string) (string, error) {
	var err error

	resp, err := http.Get("https://api.github.com/repos/" + user + "/" + repo + "/commits")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("repository %s/%s does not exist", user, repo)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	data := []CommitResponse{}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", err
	}

	if len(data) == 0 {
		return "", fmt.Errorf("repository %s/%s has no tags", user, repo)
	}

	return data[0].Sha, nil
}
