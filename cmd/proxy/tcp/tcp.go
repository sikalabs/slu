package tcp

import (
	proxy_cmd "github.com/sikalabs/slu/cmd/proxy"
	"github.com/sikalabs/slu/utils/3rdparty/go_tcp_proxy"

	"github.com/spf13/cobra"
)

var CmdFlagLocalAddr string
var CmdFlagRemoteAddr string

var Cmd = &cobra.Command{
	Use:   "tcp",
	Short: "TCP Proxy",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		go_tcp_proxy.RunProxy(CmdFlagLocalAddr, CmdFlagRemoteAddr)
	},
}

func init() {
	proxy_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&CmdFlagLocalAddr,
		"local",
		"l",
		"",
		"Local address (eg. :8000)",
	)
	Cmd.MarkFlagRequired("local")
	Cmd.Flags().StringVarP(
		&CmdFlagRemoteAddr,
		"remote",
		"r",
		"",
		"Remote address (eg. neverssl.com:80)",
	)
	Cmd.MarkFlagRequired("remote")
}
