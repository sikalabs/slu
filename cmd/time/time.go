package time

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "time",
	Short: "Time Utils",
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
