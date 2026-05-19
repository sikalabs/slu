package get_install_params

import (
	"fmt"
	"log"
	"strings"

	parentcmd "github.com/sikalabs/slu/cmd/kargo"
	"github.com/sikalabs/slu/utils/random_utils"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
)

var FlagPassword string

var Cmd = &cobra.Command{
	Use:   "get-install-params",
	Short: "Generate Kargo installation parameters (password, hashed password, signing key)",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		pass := FlagPassword
		if pass == "" {
			fmt.Print("Password (leave empty to generate): ")
			fmt.Scanln(&pass)
		}
		if pass == "" {
			pass = random_utils.RandomString(32, random_utils.ALL)
		}
		hashedBytes, err := bcrypt.GenerateFromPassword([]byte(pass), 10)
		if err != nil {
			log.Fatalln(err)
		}
		// htpasswd -B produces $2y$; Go bcrypt produces $2a$ — replace to match
		hashedPass := strings.Replace(string(hashedBytes), "$2a$", "$2y$", 1)
		signingKey := random_utils.RandomString(32, random_utils.ALL)

		fmt.Printf("Password:    %s\n", pass)
		fmt.Printf("Hashed Pass: %s\n", hashedPass)
		fmt.Printf("Signing Key: %s\n", signingKey)
	},
}

func init() {
	parentcmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagPassword,
		"password",
		"p",
		"",
		"Password to use, if not provided it will be asked (leave empty to generate)",
	)
}
