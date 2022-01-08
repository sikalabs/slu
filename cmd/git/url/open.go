package url

import (
	parent_cmd "github.com/sikalabs/slu/cmd/git"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "url",
	Short: "Get web URL or open URL of repository in Browser",
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}
