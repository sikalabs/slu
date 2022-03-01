package promdemo

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "metrics-generator",
	Short:   "Prometheus Metrics Generator Server",
	Aliases: []string{"metgen", "mg"},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
