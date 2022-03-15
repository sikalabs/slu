package ondrejsika

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "ondrejsika",
	Short:   "Ondrej Sika's Personal Utils",
	Aliases: []string{"os", "o"},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
