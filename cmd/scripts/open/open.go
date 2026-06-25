package open

import (
	scripts_cmd "github.com/sikalabs/slu/cmd/scripts"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "open",
	Short: "Open services in browser",
}

func init() {
	scripts_cmd.Cmd.AddCommand(Cmd)
}
