package wireguard

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "wireguard",
	Short:   "WireGuard Utils",
	Aliases: []string{"wg"},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
