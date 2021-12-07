package digitalocean

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "digitalocean",
	Short:   "DigitalOcean Utils",
	Aliases: []string{"do"},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
