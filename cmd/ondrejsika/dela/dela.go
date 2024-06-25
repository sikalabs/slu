package large_desktop_files

import (
	parentcmd "github.com/sikalabs/slu/cmd/ondrejsika"
	"github.com/sikalabs/slu/lib/printdela"
	"github.com/spf13/cobra"

	_ "image/jpeg"
)

var FlagRandom bool

var Cmd = &cobra.Command{
	Use:   "dela",
	Short: "Print picture of Dela to terminal",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		if FlagRandom {
			printdela.PrintRandomDela()
		} else {
			printdela.PrintDela()
		}
	},
}

func init() {
	parentcmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().BoolVarP(
		&FlagRandom,
		"random",
		"r",
		false,
		"Print random picture",
	)
}
