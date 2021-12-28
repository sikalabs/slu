package static_api

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "static-api",
	Short:   "Static API Generator (for static frontend)",
	Aliases: []string{"sapi"},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
