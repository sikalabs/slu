package latest_version

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"

	parent_cmd "github.com/sikalabs/slu/cmd/helm"
	"github.com/spf13/cobra"
)

var FlagRepo string
var FlagChart string

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

func getRepoNameFromURL(repoURL string) (string, error) {
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

var Cmd = &cobra.Command{
	Use:     "latest-version",
	Short:   "Get latest version of a Helm chart from a repository",
	Aliases: []string{"lv"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		repoName := FlagRepo

		// If FlagRepo looks like a URL, try to find the repo name
		if strings.HasPrefix(FlagRepo, "http://") || strings.HasPrefix(FlagRepo, "https://") {
			name, err := getRepoNameFromURL(FlagRepo)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
				return
			}
			repoName = name
		}

		// Execute helm search repo command with JSON output
		cmd := exec.Command("helm", "search", "repo", repoName+"/"+FlagChart, "--output", "json")
		output, err := cmd.Output()
		if err != nil {
			if exitErr, ok := err.(*exec.ExitError); ok {
				fmt.Printf("Error: %s\n", string(exitErr.Stderr))
			}
			panic(err)
		}

		// Parse JSON output
		var charts []ChartVersion
		err = json.Unmarshal(output, &charts)
		if err != nil {
			panic(err)
		}

		if len(charts) == 0 {
			fmt.Printf("No chart found for %s/%s\n", repoName, FlagChart)
			return
		}

		// Print the latest version (first result from helm search)
		fmt.Println(strings.TrimSpace(charts[0].Version))
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagRepo,
		"repo",
		"r",
		"",
		"Helm repository name or URL",
	)
	Cmd.MarkFlagRequired("repo")
	Cmd.Flags().StringVarP(
		&FlagChart,
		"chart",
		"c",
		"",
		"Chart name",
	)
	Cmd.MarkFlagRequired("chart")
}
