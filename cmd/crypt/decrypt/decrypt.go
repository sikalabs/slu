package decrypt

import (
	"fmt"
	"log"

	"github.com/sikalabs/sikalabs-crypt-go/pkg/sikalabs_crypt"
	"github.com/sikalabs/slu/cmd/crypt"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "decrypt <password> <encrypted>",
	Short: "Decrypt encrypted text with a password",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		password := args[0]
		encrypted := args[1]

		decrypted, err := sikalabs_crypt.SikaLabsSymmetricDecryptV1(password, encrypted)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(decrypted)
	},
}

func init() {
	crypt.Cmd.AddCommand(Cmd)
}
