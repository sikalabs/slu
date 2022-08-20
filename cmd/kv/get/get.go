package get

import (
	"fmt"
	"log"

	parent_cmd "github.com/sikalabs/slu/cmd/kv"
	"github.com/sikalabs/slu/lib/vault_kv"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "get",
	Short: "Get value from Key-Value store in Vault",
	Args:  cobra.ExactArgs(1),
	Run: func(c *cobra.Command, args []string) {
		key := args[0]
		value, err := vault_kv.Get(key)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(value)
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}
