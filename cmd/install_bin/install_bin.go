package install_bin

import (
	"fmt"
	"log"
	"path"
	"strings"

	"github.com/sikalabs/slu/cmd/root"
	"github.com/sikalabs/slu/utils/tar_gz_utils"
	"github.com/sikalabs/slu/utils/zip_utils"
	"github.com/spf13/cobra"
)

var CmdFlagURL string
var CmdFlagSource string
var CmdFlagName string
var CmdFlagBinDir string

var Cmd = &cobra.Command{
	Use:   "install-bin",
	Short: "Install Binary",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		if strings.HasSuffix(CmdFlagURL, "zip") {
			zip_utils.WebZipToBin(
				CmdFlagURL,
				CmdFlagSource,
				path.Join(CmdFlagBinDir, CmdFlagName),
			)
			return
		}
		if strings.HasSuffix(CmdFlagURL, "tar.gz") || strings.HasSuffix(CmdFlagURL, "tgz") {
			tar_gz_utils.WebTarGzToBin(
				CmdFlagURL,
				CmdFlagSource,
				path.Join(CmdFlagBinDir, CmdFlagName),
			)
			return
		}
		log.Fatal(fmt.Errorf("unknown suffix (no .zip, .tar.gz or .tgz)"))
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
		"Source path in archive (zip, tar.gz)",
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
