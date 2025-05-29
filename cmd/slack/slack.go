package slack

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "slack",
	Short: "Slack Utils",
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
