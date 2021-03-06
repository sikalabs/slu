package install_bin

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/sikalabs/slu/utils/install_bin_utils"
	"github.com/spf13/cobra"
)

var CmdFlagURL string
var CmdFlagSource string
var CmdFlagName string
var CmdFlagBinDir string

var Cmd = &cobra.Command{
	Use:     "install-bin",
	Short:   "Install Binary",
	Aliases: []string{"ib"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		source := CmdFlagName
		if CmdFlagSource != "" {
			source = CmdFlagSource
		}
		install_bin_utils.InstallBin(
			CmdFlagURL,
			source,
			CmdFlagBinDir,
			CmdFlagName,
		)
	},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&CmdFlagURL,
		"url",
		"u",
		"",
		"Url to source (bin, zip, tar.gz)",
	)
	Cmd.MarkFlagRequired("url")
	Cmd.Flags().StringVarP(
		&CmdFlagSource,
		"source",
		"s",
		"",
		"Source path in archive (default from --name)",
	)
	Cmd.Flags().StringVarP(
		&CmdFlagName,
		"name",
		"n",
		"",
		"Name of binary",
	)
	Cmd.MarkFlagRequired("name")
	Cmd.Flags().StringVarP(
		&CmdFlagBinDir,
		"bin-dir",
		"d",
		"/usr/local/bin",
		"Binary dir",
	)
}
