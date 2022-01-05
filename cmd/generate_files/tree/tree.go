package tree

import (
	parent_cmd "github.com/sikalabs/slu/cmd/generate_files"
	"github.com/sikalabs/slu/utils/generate_files_utils"

	"github.com/spf13/cobra"
)

var FlagPath string
var FlagSizeMB int
var FlagCount int

var Cmd = &cobra.Command{
	Use:   "tree",
	Short: "Generate tree of files",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		generate_files_utils.GenerateFiles(FlagPath, FlagSizeMB, FlagCount)
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagPath,
		"path",
		"p",
		"",
		"Target path",
	)
	Cmd.MarkFlagRequired("path")
	Cmd.Flags().IntVarP(
		&FlagSizeMB,
		"size",
		"s",
		1,
		"Size in MB",
	)
	Cmd.Flags().IntVarP(
		&FlagCount,
		"count",
		"c",
		0,
		"Count of new files",
	)
	Cmd.MarkFlagRequired("count")
}
