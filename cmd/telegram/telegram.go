package telegram

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "telegram",
	Short: "Telegram Utils",
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
