package rename

import (
	parent_cmd "github.com/sikalabs/slu/cmd/scripts"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "rename",
	Short: "Rename files",
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}
