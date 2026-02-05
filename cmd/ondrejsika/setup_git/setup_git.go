package clear_dns_cache_mac

import (
	parentcmd "github.com/sikalabs/slu/cmd/ondrejsika"
	"github.com/sikalabs/slu/utils/exec_utils"
	"github.com/spf13/cobra"
)

var FlagAsDela bool

var Cmd = &cobra.Command{
	Use:   "setup-git",
	Short: "Setup Git as Ondrej Sika",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		name := "Ondrej Sika"
		email := "ondrej@ondrejsika.com"
		if FlagAsDela {
			name = "Dela"
			email = "dela@sikalabs.com"
		}
		exec_utils.ExecOut(
			"git", "config", "--global", "user.name",
			name)
		exec_utils.ExecOut(
			"git", "config", "--global", "user.email",
			email)
		exec_utils.ExecOut(
			"git", "config", "--global", "core.editor",
			"vim")
	},
}

func init() {
	parentcmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().BoolVar(&FlagAsDela, "as-dela", false, "Setup Git as Dela (dela@sikalabs.com)")
}
