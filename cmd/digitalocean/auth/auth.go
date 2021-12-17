package auth

import (
	parent_cmd "github.com/sikalabs/slu/cmd/digitalocean"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "auth",
	Short: "slu authentification to DigitalOcean",
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}
