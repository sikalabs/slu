package password

import (
	"fmt"
	"log"

	random_cmd "github.com/sikalabs/slu/cmd/random"
	"github.com/sikalabs/slu/utils/random_utils"

	"github.com/spf13/cobra"
)

var FlagDouble bool

var Cmd = &cobra.Command{
	Use:     "password",
	Short:   "Generate random password",
	Aliases: []string{"pwd", "passwd", "pass"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		password := getPasswordOrDie()
		if FlagDouble {
			password = password + `_` + getPasswordOrDie()
		}
		fmt.Println(password)
	},
}

func init() {
	random_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().BoolVar(
		&FlagDouble,
		"double",
		false,
		"Double the length of the passowd",
	)
}

func getPasswordOrDie() string {
	password, err := random_utils.RandomPassword()
	if err != nil {
		log.Fatal(err)
	}
	return password
}
