package helm_utils

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

type ChartVersion struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	AppVersion  string `json:"app_version"`
	Description string `json:"description"`
}

type RepoEntry struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type ChartMetadata struct {
	Version string `yaml:"version"`
}

func GetRepoNameFromURL(repoURL string) (string, error) {
	// Normalize URL - remove trailing slash
	repoURL = strings.TrimSuffix(repoURL, "/")

	// Get list of repos
	cmd := exec.Command("helm", "repo", "list", "--output", "json")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	var repos []RepoEntry
	err = json.Unmarshal(output, &repos)
	if err != nil {
		return "", err
	}

	// Find repo by URL
	for _, repo := range repos {
		if strings.TrimSuffix(repo.URL, "/") == repoURL {
			return repo.Name, nil
		}
	}

	return "", fmt.Errorf("no repository found with URL: %s", repoURL)
}

func GetLatestVersionFromOCI(ociURL string) (string, error) {
	// Execute helm show chart command
	cmd := exec.Command("helm", "show", "chart", ociURL)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	// Extract version from the "Pulled:" line using regex
	pulledRegex := regexp.MustCompile(`Pulled: [^:]+:([^\s]+)`)
	matches := pulledRegex.FindStringSubmatch(string(output))
	if len(matches) > 1 {
		return matches[1], nil
	}

	// Fallback: parse YAML to get version
	var metadata ChartMetadata
	err = yaml.Unmarshal(output, &metadata)
	if err != nil {
		return "", err
	}

	return metadata.Version, nil
}

func GetLatestVersionFromRepo(repoName, chartName string) (string, error) {
	// Execute helm search repo command with JSON output
	cmd := exec.Command("helm", "search", "repo", repoName+"/"+chartName, "--output", "json")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	// Parse JSON output
	var charts []ChartVersion
	err = json.Unmarshal(output, &charts)
	if err != nil {
		return "", err
	}

	if len(charts) == 0 {
		return "", fmt.Errorf("no chart found for %s/%s", repoName, chartName)
	}

	// Return the latest version (first result from helm search)
	return strings.TrimSpace(charts[0].Version), nil
}
