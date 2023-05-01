package atol

import (
	parent_cmd "github.com/sikalabs/slu/cmd/scripts"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "atol",
	Short: "atol scripts",
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}
