package version

import (
	"fmt"
	"runtime"

	"github.com/sikalabs/slu/cmd/root"
	"github.com/sikalabs/slu/version"
	"github.com/spf13/cobra"
)

var CmdFlagVerbose bool

var Cmd = &cobra.Command{
	Use:     "version",
	Short:   "Prints version",
	Aliases: []string{"v"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		if CmdFlagVerbose {
			fmt.Printf("%s %s %s\n", version.Version, runtime.GOOS, runtime.GOARCH)
		} else {
			fmt.Printf("%s\n", version.Version)
		}
	},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
	Cmd.Flags().BoolVarP(
		&CmdFlagVerbose,
		"verbose",
		"v",
		false,
		"Verbose version output",
	)
}
