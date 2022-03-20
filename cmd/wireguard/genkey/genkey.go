package genkey

import (
	parent_cmd "github.com/sikalabs/slu/cmd/wireguard"
	"github.com/sikalabs/slu/utils/wireguard_utils"
	"github.com/spf13/cobra"
)

var FlagJson bool
var CmdFlagAddr string
var CmdFlagTimeout int

var Cmd = &cobra.Command{
	Use:     "genkey",
	Short:   "Genetate WireGuard private, public & preshared keys",
	Aliases: []string{"gen", "g", "key", "keygen"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		if FlagJson {
			wireguard_utils.PrintNewKeysJson()
		} else {
			wireguard_utils.PrintNewKeys()
		}
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
	Cmd.PersistentFlags().BoolVar(
		&FlagJson,
		"json",
		false,
		"Format output to JSON",
	)
}
