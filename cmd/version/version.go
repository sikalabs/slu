package version

import (
	"fmt"

	"github.com/sikalabs/slut/cmd/root"
	"github.com/sikalabs/slut/version"
	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Prints version",
	Aliases: []string{"v"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		fmt.Printf("%s\n", version.Version)
	},
}

func init() {
	root.RootCmd.AddCommand(VersionCmd)
}
