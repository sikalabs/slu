package server

import (
	"strconv"

	parent_cmd "github.com/sikalabs/slu/cmd/metrics_generator"
	libserver "github.com/sikalabs/slu/lib/metrics_generator/server"

	"github.com/spf13/cobra"
)

var FlagPort int

var Cmd = &cobra.Command{
	Use:     "server",
	Short:   "Run Prometheus Demo Server",
	Aliases: []string{"s"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		addr := ":" + strconv.Itoa(FlagPort)
		libserver.Server(addr)
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().IntVarP(
		&FlagPort,
		"port",
		"p",
		8000,
		"listen on port",
	)
}
