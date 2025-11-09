package create

import (
	"log"

	parent_cmd "github.com/sikalabs/slu/cmd/vaultino"
	vaultinoUtils "github.com/sikalabs/slu/utils/vaultino"
	"github.com/spf13/cobra"
)

var (
	FlagFile string
)

var Cmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"crt"},
	Short:   "Create new Vaultino encrypted file",
	Long:    "Create new Vaultino encrypted file. It will prompt for a password and create an encrypted file.",
	Example: "vp-utils vaultino create <name> --file <path/to/encrypted_file>",
	Run: func(cmd *cobra.Command, args []string) {
		if args == nil || len(args) < 1 {
			log.Fatalf("name of the encrypted file is required as the first argument")
		}
		err := vaultinoUtils.CreateVault(args[0], FlagFile)
		if err != nil {
			log.Fatalf("failed to create vault: %v", err)
		}
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(&FlagFile, "file", "f", "", "Path to source file to encrypt")
	Cmd.MarkFlagRequired("file")
}
