package validate

import (
	"log"

	parent_cmd "github.com/sikalabs/slu/cmd/jwt"
	"github.com/sikalabs/slu/utils/jwt_utils"

	"github.com/spf13/cobra"
)

var FlagVerbose bool

var Cmd = &cobra.Command{
	Use:   "validate <issuer> <rawToken>",
	Short: "Validate JWT",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		err := jwt_utils.ValidateJWT(args[0], args[1])
		if err != nil {
			log.Fatalln(err)
		}
		if FlagVerbose {
			log.Println("JWT is valid")
		}
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().BoolVarP(
		&FlagVerbose,
		"verbose",
		"v",
		false,
		"Verbose output",
	)
}
