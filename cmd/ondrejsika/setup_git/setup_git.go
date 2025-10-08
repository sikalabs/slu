package clear_dns_cache_mac

import (
	parentcmd "github.com/sikalabs/slu/cmd/ondrejsika"
	"github.com/sikalabs/slu/utils/exec_utils"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "setup-git",
	Short: "Setup Git as Ondrej Sika",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		exec_utils.ExecOut(
			"git", "config", "--global", "user.name",
			"Ondrej Sika")
		exec_utils.ExecOut(
			"git", "config", "--global", "user.email",
			"ondrej@ondrejsika.com")
		exec_utils.ExecOut(
			"git", "config", "--global", "core.editor",
			"vim")
	},
}

func init() {
	parentcmd.Cmd.AddCommand(Cmd)
}
