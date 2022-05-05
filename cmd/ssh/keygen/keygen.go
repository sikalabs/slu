package keygen

import (
	"fmt"
	"log"

	parent_cmd "github.com/sikalabs/slu/cmd/ssh"
	"github.com/sikalabs/slu/utils/ssh_utils"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "keygen",
	Short:   "Genetate SSH private key",
	Aliases: []string{"gen", "g", "key", "keygen"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		pub, priv, err := ssh_utils.MakeSSHKeyPair()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(priv)
		fmt.Println(pub)
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}
