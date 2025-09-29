package password_hash

import (
	"fmt"
	"log"

	parentcmd "github.com/sikalabs/slu/cmd/argocd"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
)

var FlagPassword string

var Cmd = &cobra.Command{
	Use:   "password-hash",
	Short: "Create a bcrypt password hash for ArgoCD",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		printPasswordHash(FlagPassword)
	},
}

func init() {
	parentcmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagPassword,
		"password",
		"p",
		"",
		"Password to hash, if not provided, it will be asked",
	)
}

func printPasswordHash(password string) {
	if password == "" {
		fmt.Print("Password: ")
		fmt.Scanln(&password)
	}

	// Generate the bcrypt hash of the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(hashedPassword))
}