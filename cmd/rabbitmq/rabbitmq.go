package rabbitmq

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "rabbitmq",
	Aliases: []string{"rmq"},
	Short:   "RabbitMQ CLI Utils",
	Long:    "Custom RabbitMQ CLI Utils, which can be used for testing and managing RabbitMQ servers.",
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
