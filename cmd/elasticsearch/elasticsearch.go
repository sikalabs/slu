package elasticsearch

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "elasticsearch",
	Short:   "ElasticSearch Utils",
	Aliases: []string{"es"},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
