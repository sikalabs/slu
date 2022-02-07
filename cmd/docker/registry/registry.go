package all

import (
	parent_cmd "github.com/sikalabs/slu/cmd/docker"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "registry",
	Short:   "Docker Registry Utils",
	Aliases: []string{"reg", "r"},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}
