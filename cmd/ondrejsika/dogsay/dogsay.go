package dogsay

import (
	say "github.com/sikalabs/dogsay/pkg/dogsay"
	parentcmd "github.com/sikalabs/slu/cmd/ondrejsika"
	"github.com/spf13/cobra"

	_ "image/jpeg"
)

var FlagRandom bool

var Cmd = &cobra.Command{
	Use:   "dogsay <text>",
	Short: "Like cowsay but with doggo",
	Args:  cobra.ExactArgs(1),
	Run: func(c *cobra.Command, args []string) {
		say.PrintDogSay(args[0])
	},
}

func init() {
	parentcmd.Cmd.AddCommand(Cmd)
}
