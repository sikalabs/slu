package http200

import (
	parentcmd "github.com/sikalabs/slu/cmd/wait_for"
	"github.com/sikalabs/slu/utils/wait_for_http200_utils"
	"github.com/spf13/cobra"
)

var CmdFlagUrl string
var CmdFlagTimeout int

var Cmd = &cobra.Command{
	Use:   "http200",
	Short: "Wait for HTTP 200 response",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		wait_for_http200_utils.WaitForHttp200(CmdFlagTimeout, CmdFlagUrl)
	},
}

func init() {
	parentcmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&CmdFlagUrl,
		"url",
		"u",
		"",
		"URL to check (eg.: http://example.com)",
	)
	Cmd.MarkFlagRequired("url")
	Cmd.Flags().IntVarP(
		&CmdFlagTimeout,
		"timeout",
		"t",
		5*60, // 5 min
		"Timeout in seconds",
	)
}
