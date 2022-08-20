package keygen

import (
	"fmt"
	"log"

	parent_cmd "github.com/sikalabs/slu/cmd/ssh"
	"github.com/sikalabs/slu/utils/ssh_utils"
	"github.com/spf13/cobra"
)

const DEFAULT_KEY_LENGTH = 2048

var FlagUseGo bool
var FlagECDSA bool

var Cmd = &cobra.Command{
	Use:     "keygen",
	Short:   "Genetate SSH private key",
	Aliases: []string{"gen", "g", "key", "keygen"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		var pub string
		var priv string
		var err error
		length := DEFAULT_KEY_LENGTH
		if FlagUseGo {
			// Go SSH key generation
			pub, priv, err = ssh_utils.MakeSSHKeyPair()
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println(priv)
			fmt.Println(pub)
		} else {
			// OpenSSH key generation
			if FlagECDSA {
				pub, priv, err = ssh_utils.MakeSSHKeyPairSSHKeyGenECDSA()
			} else {
				pub, priv, err = ssh_utils.MakeSSHKeyPairSSHKeyGen(length)
			}
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println(priv)
			fmt.Println(pub)
		}
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
	Cmd.PersistentFlags().BoolVar(
		&FlagUseGo,
		"use-go",
		false,
		"Use Go implementation of SSH key generation",
	)
	Cmd.PersistentFlags().BoolVar(
		&FlagECDSA,
		"ecdsa",
		false,
		"Use ECDSA key generation",
	)
}
