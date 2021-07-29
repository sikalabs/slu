package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version string = "v0.1.0-dev"

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints version",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		fmt.Printf("%s\n", version)
	},
}
