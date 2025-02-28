package password

import (
	parent_cmd "github.com/sikalabs/slu/cmd/jwt"
	"github.com/sikalabs/slu/utils/jwt_utils"
	"github.com/sikalabs/slu/utils/stdin_utils"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "parse <jwt>",
	Short: "Parse JWT from stdin into JSON list of 3 objects (Header, Payload, Signature)",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		jwtToken := args[0]
		if jwtToken == "-" {
			jwtToken = stdin_utils.ReadFromPipeOrDie()
		}
		jwt_utils.ParseJWT(jwtToken)
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}
