package version

import (
	"os"

	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var (
	flagOutputDir string
)

var Cmd = &cobra.Command{
	Use:   "generate-docs",
	Short: "Generate Markdown docs",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		path := flagOutputDir
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			panic(err)
		}
		err = doc.GenMarkdownTree(root.RootCmd, path)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	Cmd.Flags().StringVarP(
		&flagOutputDir,
		"output-dir",
		"o",
		"./cobra-docs/",
		"Output directory for generated docs",
	)
	root.RootCmd.AddCommand(Cmd)
}
