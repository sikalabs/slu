package go_code

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "go-code",
	Short: "Utils for writing Go code",
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
