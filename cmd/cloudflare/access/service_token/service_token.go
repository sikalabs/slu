package service_token

import (
	parent_cmd "github.com/sikalabs/slu/cmd/cloudflare/access"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "service-token",
	Short:   "Cloudflare Access Service Token Utils",
	Aliases: []string{"st"},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}
