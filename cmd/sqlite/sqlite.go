package sqlite

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "sqlite",
	Short: "SQLite Utils",
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
