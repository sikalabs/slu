package docker

import (
	parent_cmd "github.com/sikalabs/slu/cmd/scripts"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "docker",
	Short:   "Docker Scripts",
	Aliases: []string{"d"},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}
