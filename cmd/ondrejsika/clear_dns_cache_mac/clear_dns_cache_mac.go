package clear_dns_cache_mac

import (
	parentcmd "github.com/sikalabs/slu/cmd/ondrejsika"
	"github.com/sikalabs/slu/utils/exec_utils"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "clear-dns-cache-mac",
	Short: "Clear DNS Cache on Mac",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		exec_utils.ExecOut("killall", "-HUP", "mDNSResponder")
	},
}

func init() {
	parentcmd.Cmd.AddCommand(Cmd)
}
