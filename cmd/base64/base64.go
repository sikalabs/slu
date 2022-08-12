package base64

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "base64",
	Short:   "Base64 Utils",
	Aliases: []string{"b64"},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
