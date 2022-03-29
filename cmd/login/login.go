package login

import (
	"fmt"
	"log"
	"syscall"

	"github.com/sikalabs/slu/cmd/root"
	"github.com/sikalabs/slu/config"
	"github.com/sikalabs/slu/utils/vault_utils"
	"golang.org/x/crypto/ssh/terminal"

	"github.com/spf13/cobra"
)

var FlagUrl string
var FlagUsename string
var FlagPassword string

var Cmd = &cobra.Command{
	Use:   "login",
	Short: "Login to slu vault",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		client, err := vault_utils.GetClient(FlagUrl)
		if err != nil {
			log.Fatalln(err)
		}
		password := FlagPassword
		if password == "" {
			fmt.Print("Password: ")
			passwordBytes, err := terminal.ReadPassword(int(syscall.Stdin))
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Print("\n")
			password = string(passwordBytes)
		}
		token, err := vault_utils.GetTokenFromUserpass(client, FlagUsename, password)
		if err != nil {
			log.Fatalln(err)
		}
		co := config.ReadConfig()
		co.SluVault.Url = FlagUrl
		co.SluVault.User = FlagUsename
		config.WriteConfig(co)
		se := config.ReadSecrets()
		se.SluVault.Token = token
		config.WriteSecrets(se)
		fmt.Println("Success!")
	},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagUrl,
		"url",
		"U",
		"",
		"Vault URL",
	)
	Cmd.MarkFlagRequired("url")
	Cmd.Flags().StringVarP(
		&FlagUsename,
		"user",
		"u",
		"",
		"Vault username",
	)
	Cmd.MarkFlagRequired("user")
	Cmd.Flags().StringVarP(
		&FlagPassword,
		"password",
		"p",
		"",
		"Vault password",
	)
}
