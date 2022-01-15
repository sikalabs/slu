package du

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/sikalabs/slu/utils/du_utils"
	"github.com/spf13/cobra"
)

var FlagNoHumanReadable bool
var FlagThreshold string
var FlagMaxDepth int

var Cmd = &cobra.Command{
	Use:   "du",
	Short: "Own implemetation of \"du\"",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		du_utils.RunDiskUsage(!FlagNoHumanReadable, FlagThreshold, FlagMaxDepth)
	},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagThreshold,
		"threshold",
		"t",
		"",
		"threshold of the size, any folders' size larger than the threshold will be print. for example, '1G', '10M', '100K', '1024'",
	)
	Cmd.Flags().BoolVarP(
		&FlagNoHumanReadable,
		"human-readable",
		"H",
		false,
		"disable human readable unit of size",
	)
	Cmd.Flags().IntVarP(
		&FlagMaxDepth,
		"max-depth",
		"d",
		0,
		"list its subdirectories and their sizes to any desired level of depth (i.e., to any level of subdirectories) in a directory tree.",
	)
}
