package password

import (
	"encoding/json"
	"fmt"

	parent_cmd "github.com/sikalabs/slu/cmd/awx"
	"github.com/sikalabs/slu/utils/awx_utils"

	"github.com/spf13/cobra"
)

var FlagJson bool
var FlagNamespace string
var FlagAWX string

var Cmd = &cobra.Command{
	Use:     "password",
	Short:   "Get AWX Admin Passowrd",
	Aliases: []string{"p"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		password := awx_utils.GetAWXPassword(FlagNamespace, FlagAWX)

		if FlagJson {
			outJson, err := json.Marshal(password)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(outJson))
		} else {
			fmt.Println(password)
		}
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagNamespace,
		"namespace",
		"n",
		"awx",
		"AWX Namespace",
	)
	Cmd.Flags().StringVarP(
		&FlagAWX,
		"awx",
		"a",
		"awx",
		"AWX Instance namw",
	)
	Cmd.PersistentFlags().BoolVar(
		&FlagJson,
		"json",
		false,
		"Format output to JSON",
	)
}
