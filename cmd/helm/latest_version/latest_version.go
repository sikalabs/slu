package latest_version

import (
	"fmt"
	"log"
	"strings"

	parent_cmd "github.com/sikalabs/slu/cmd/helm"
	"github.com/sikalabs/slu/utils/helm_utils"
	"github.com/spf13/cobra"
)

var FlagRepo string
var FlagChart string

var Cmd = &cobra.Command{
	Use:     "latest-version",
	Short:   "Get latest version of a Helm chart from a repository",
	Aliases: []string{"lv"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		// Handle OCI registries
		if strings.HasPrefix(FlagRepo, "oci://") {
			ociURL := FlagRepo + "/" + FlagChart
			version, err := helm_utils.GetLatestVersionFromOCI(ociURL)
			handleError(err, "Error getting latest version from OCI registry")
			fmt.Println(version)
			return
		}

		repoName := FlagRepo

		// If FlagRepo looks like a URL, try to find the repo name
		if strings.HasPrefix(FlagRepo, "http://") || strings.HasPrefix(FlagRepo, "https://") {
			name, err := helm_utils.GetRepoNameFromURL(FlagRepo)
			handleError(err, "Error finding repository by URL")
			repoName = name
		}

		// Get latest version from repo
		version, err := helm_utils.GetLatestVersionFromRepo(repoName, FlagChart)
		handleError(err, "Error getting latest version from repository")

		fmt.Println(version)
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagRepo,
		"repo",
		"r",
		"",
		"Helm repository name, URL, or OCI registry URL (oci://...)",
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

func handleError(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %v\n", message, err)
	}
}
