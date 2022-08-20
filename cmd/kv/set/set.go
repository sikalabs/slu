package set

import (
	"log"

	parent_cmd "github.com/sikalabs/slu/cmd/kv"
	"github.com/sikalabs/slu/lib/vault_kv"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "set",
	Short: "Set value from Key-Value store in Vault",
	Args:  cobra.ExactArgs(2),
	Run: func(c *cobra.Command, args []string) {
		key := args[0]
		value := args[1]
		err := vault_kv.Set(key, value)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}
