package all

import (
	parent_cmd "github.com/sikalabs/slu/cmd/digitalocean"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "all",
	Short: "Work with all resources in DigitalOcean account",
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}
