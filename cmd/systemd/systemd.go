package systemd

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "systemd",
	Short:   "Utils for systemd",
	Aliases: []string{"sysd"},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
