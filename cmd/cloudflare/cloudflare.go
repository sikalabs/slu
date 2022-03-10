package cloudflare

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "cloudflare",
	Short:   "Cloudflare Utils",
	Aliases: []string{"cf"},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
