package genkey

import (
	root_cmd "github.com/sikalabs/slu/cmd/root"
	parent_cmd "github.com/sikalabs/slu/cmd/wireguard"
	"github.com/sikalabs/slu/utils/wireguard_utils"
	"github.com/spf13/cobra"
)

var CmdFlagAddr string
var CmdFlagTimeout int

var Cmd = &cobra.Command{
	Use:     "genkey",
	Short:   "Genetate WireGuard private, public & preshared keys",
	Aliases: []string{"gen", "g", "key", "keygen"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		if root_cmd.RootCmdFlagJson {
			wireguard_utils.PrintNewKeysJson()
		} else {
			wireguard_utils.PrintNewKeys()
		}
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}
