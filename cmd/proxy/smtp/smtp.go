package smtp

import (
	proxy_cmd "github.com/sikalabs/slu/cmd/proxy"
	"github.com/sikalabs/slu/utils/smtp_proxy_utils"

	"github.com/spf13/cobra"
)

var CmdFlagLocalAddr string
var CmdFlagRemoteAddr string

var Cmd = &cobra.Command{
	Use:   "smtp",
	Short: "SMTP Proxy",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		smtp_proxy_utils.RunSimpleSMTPProxy(CmdFlagLocalAddr, CmdFlagRemoteAddr)
	},
}

func init() {
	proxy_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&CmdFlagLocalAddr,
		"local",
		"l",
		"",
		"Local address (eg. :1025)",
	)
	Cmd.MarkFlagRequired("local")
	Cmd.Flags().StringVarP(
		&CmdFlagRemoteAddr,
		"remote",
		"r",
		"",
		"Remote address (eg. smtp.com:25)",
	)
	Cmd.MarkFlagRequired("remote")
}
