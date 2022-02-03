package helm

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "helm",
	Short: "Helm Utils",
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
