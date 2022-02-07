package digitalocean

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "docker",
	Short: "Docker Utils",
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
