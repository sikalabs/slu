package from_base64

import (
	"log"

	parent_cmd "github.com/sikalabs/slu/cmd/file_utils"
	"github.com/sikalabs/slu/utils/file_utils"

	"github.com/spf13/cobra"
)

var FlagPath string
var FlagContent string

var Cmd = &cobra.Command{
	Use:     "from-base64",
	Short:   "Create file from base64 encoded string",
	Aliases: []string{"fb64"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		err := file_utils.CreateFileFromBase64(FlagPath, FlagContent)
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
		&FlagContent,
		"content",
		"c",
		"",
		"Base64 encoded content",
	)
	Cmd.MarkFlagRequired("content")
}
