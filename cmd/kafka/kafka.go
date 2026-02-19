package kafka

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "kafka",
	Short: "Kafka CLI Utils",
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
