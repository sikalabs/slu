package download

import (
	"log"

	parent_cmd "github.com/sikalabs/slu/cmd/file_utils"
	"github.com/sikalabs/slu/utils/file_utils"

	"github.com/spf13/cobra"
)

var FlagPath string
var FlagUrl string

var Cmd = &cobra.Command{
	Use:     "download",
	Short:   "Download file from URL",
	Aliases: []string{"d"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		err := file_utils.DownloadFile(FlagPath, FlagUrl)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagPath,
		"path",
		"p",
		"",
		"Output file path",
	)
	Cmd.MarkFlagRequired("path")
	Cmd.Flags().StringVarP(
		&FlagUrl,
		"url",
		"u",
		"",
		"File URL",
	)
	Cmd.MarkFlagRequired("url")
}
