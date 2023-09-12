package sentry

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "sentry",
	Short: "Sentry utils",
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
