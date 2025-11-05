package encrypt

import (
	"fmt"
	"log"

	"github.com/sikalabs/sikalabs-crypt-go/pkg/sikalabs_crypt"
	"github.com/sikalabs/slu/cmd/crypt"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "encrypt <password> <text>",
	Short: "Encrypt text with a password",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		password := args[0]
		text := args[1]

		encrypted, err := sikalabs_crypt.SikaLabsSymmetricEncryptV1(password, text)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(encrypted)
	},
}

func init() {
	crypt.Cmd.AddCommand(Cmd)
}
