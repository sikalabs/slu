package use_ssh

import (
	parent_cmd "github.com/sikalabs/slu/cmd/git"
	"github.com/sikalabs/slu/utils/git_utils"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "use-ssh",
	Short: "Switch from HTTPS to SSH",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		git_utils.UseSSH()
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}
