package edit

import (
	"log"

	parent_cmd "github.com/sikalabs/slu/cmd/vaultino"
	vaultinoUtils "github.com/sikalabs/slu/utils/vaultino"

	"github.com/spf13/cobra"
)

var FlagChangePassword bool

var Cmd = &cobra.Command{
	Use:     "edit",
	Aliases: []string{"e"},
	Short:   "Edit Vaultino encrypted file",
	Long:    "Edit Vaultino encrypted file. It will prompt for a password, decrypt the file, open it in the default editor, and re-encrypt it upon saving.",
	Example: "vp-utils vaultino edit <path_tp_file>\nvp-utils vaultino edit --change-password <path_tp_file>",
	Run: func(cmd *cobra.Command, args []string) {
		if args == nil || len(args) < 1 {
			log.Fatalf("path to the encrypted file is required as the first argument")
		}

		var err error
		if FlagChangePassword {
			err = vaultinoUtils.EditVaultWithPasswordChange(args[0])
		} else {
			err = vaultinoUtils.EditVault(args[0])
		}

		if err != nil {
			log.Fatalf("failed to edit vault: %v", err)
		}
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().BoolVarP(&FlagChangePassword, "change-password", "p", false, "Change the vault password after editing")
}
