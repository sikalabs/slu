package iceland

import (
	parentcmd "github.com/sikalabs/slu/cmd/ondrejsika"
	"github.com/sikalabs/slu/lib/printiceland"
	"github.com/spf13/cobra"

	_ "image/jpeg"
)

var Cmd = &cobra.Command{
	Use:   "iceland",
	Short: "Print picture from Iceland to terminal",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		printiceland.PrintRadomIcelandPhoto()
	},
}

func init() {
	parentcmd.Cmd.AddCommand(Cmd)
}
