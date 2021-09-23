package sleep

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "sleep",
	Short: "Sleep Utils",
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
