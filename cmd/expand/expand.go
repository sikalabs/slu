package expand

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "expand",
	Short: "Expand environment variables in files and strings",
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
