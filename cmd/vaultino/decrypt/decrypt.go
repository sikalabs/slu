package decrypt

import (
	"log"

	parent_cmd "github.com/sikalabs/slu/cmd/vaultino"
	vaultinoUtils "github.com/sikalabs/slu/utils/vaultino"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "decrypt",
	Aliases: []string{"crt"},
	Short:   "Decrypt Vaultino encrypted file",
	Long:    "Decrypt Vaultino encrypted file. It will prompt for a password and create a decrypted file.",
	Example: "vp-utils vaultino decrypt <path_tp_file>",
	Run: func(cmd *cobra.Command, args []string) {
		if args == nil || len(args) < 1 {
			log.Fatalf("path to the encrypted file is required as the first argument")
		}
		err := vaultinoUtils.DecryptVaultToFile(args[0])
		if err != nil {
			log.Fatalf("failed to decrypt vault: %v", err)
		}
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}
