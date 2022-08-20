package host

import (
	parent_cmd "github.com/sikalabs/slu/cmd/cloudflare"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "access",
	Short:   "Cloudflare Access Utils",
	Aliases: []string{"a"},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}
