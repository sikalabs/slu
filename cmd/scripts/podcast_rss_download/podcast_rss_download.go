package podcast_rss_download

import (
	parent_cmd "github.com/sikalabs/slu/cmd/scripts"
	"github.com/sikalabs/slu/utils/podcast_rss_download_utils"
	"github.com/spf13/cobra"
)

var FlagRssUrl string
var FlagOutDir string

var Cmd = &cobra.Command{
	Use:     "podcast-rss-download",
	Aliases: []string{"prssd"},
	Short:   "Download podcast episodes from RSS",
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		podcast_rss_download_utils.PodcastRssDownload(FlagRssUrl, FlagOutDir)
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagRssUrl,
		"rss-url",
		"u",
		"",
		"URL of the podcast RSS feed",
	)
	Cmd.MarkFlagRequired("rss-url")
	Cmd.Flags().StringVarP(
		&FlagRssUrl,
		"out-dir",
		"d",
		".",
		"Output directory for MP3 files",
	)
}
