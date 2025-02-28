package jwt

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "jwt",
	Short: "JWT Utils",
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
