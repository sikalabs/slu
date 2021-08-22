package parse_file

import (
	tls_cmd "github.com/sikalabs/slu/cmd/tls"
	"github.com/sikalabs/slu/utils/tls_utils"

	"github.com/spf13/cobra"
)

var CmdFlagCertFile string
var CmdFlagKeyFile string

var Cmd = &cobra.Command{
	Use:   "parse-file",
	Short: "Parse TLS Certificate data from file (cert, key)",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		tls_utils.PrintCertificateFromFile(CmdFlagCertFile, CmdFlagKeyFile)
	},
}

func init() {
	tls_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&CmdFlagCertFile,
		"cert",
		"c",
		"",
		"Cert file",
	)
	Cmd.MarkFlagRequired("cert")
	Cmd.Flags().StringVarP(
		&CmdFlagKeyFile,
		"key",
		"k",
		"",
		"Key file",
	)
	Cmd.MarkFlagRequired("key")

}
