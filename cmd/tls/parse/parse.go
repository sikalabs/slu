package parse

import (
	tls_cmd "github.com/sikalabs/slu/cmd/tls"
	"github.com/sikalabs/slu/utils/tls_utils"

	"github.com/spf13/cobra"
)

var CmdFlagAddr string
var CmdFlagServerName string

var Cmd = &cobra.Command{
	Use:   "parse",
	Short: "Parse TLS Certificate data from server (connection)",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		tls_utils.PrintCertificateFromServer(CmdFlagAddr, CmdFlagServerName)
	},
}

func init() {
	tls_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&CmdFlagAddr,
		"address",
		"a",
		"",
		"Address (eg.: google.com:443)",
	)
	Cmd.MarkFlagRequired("address")
	Cmd.Flags().StringVarP(
		&CmdFlagServerName,
		"server-name",
		"n",
		"",
		"ServerName (SNI)",
	)
}
