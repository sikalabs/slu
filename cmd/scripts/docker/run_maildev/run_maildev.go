package run_maildev

import (
	"fmt"

	parent_cmd "github.com/sikalabs/slu/cmd/scripts/docker"
	"github.com/sikalabs/slu/utils/exec_utils"
	"github.com/spf13/cobra"
)

var FlagDry bool
var FlagNamespace string
var FlagDomain string
var FlagInstallOnly bool

var Cmd = &cobra.Command{
	Use:   "run-maildev",
	Short: "docker run --name maildev -d -p 1080:1080 -p 1025:1025 maildev/maildev",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		exec_utils.ExecOut(
			"docker",
			"run",
			"--name", "maildev",
			"-d",
			"-p", "1080:1080",
			"-p", "1025:1025",
			"maildev/maildev",
		)
		fmt.Println("See: http://127.0.0.1:1080")
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}
