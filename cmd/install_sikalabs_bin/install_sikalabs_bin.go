package install_sikalabs_bin

import (
	"fmt"
	"runtime"

	"github.com/sikalabs/slu/cmd/root"
	"github.com/sikalabs/slu/utils/install_bin_utils"
	"github.com/spf13/cobra"
)

var CmdFlagOS string
var CmdFlagArch string
var CmdFlagBinDir string

var Cmd = &cobra.Command{
	Use:     "install-sikalabs-bin <name>",
	Short:   "Install binary from SikaLabs bin bucket (sikalabs.fra1.cdn.digitaloceanspaces.com/bin)",
	Aliases: []string{"islb"},
	Args:    cobra.ExactArgs(1),
	Run: func(c *cobra.Command, args []string) {
		name := args[0]
		url := fmt.Sprintf(
			"https://sikalabs.fra1.cdn.digitaloceanspaces.com/bin/%s-%s-%s",
			name,
			CmdFlagOS,
			CmdFlagArch,
		)
		install_bin_utils.InstallBin(
			url,
			name,
			CmdFlagBinDir,
			name,
			CmdFlagOS == "windows",
		)
	},
}

func init() {
	defaultBinDir := "/usr/local/bin"
	if runtime.GOOS == "windows" {
		defaultBinDir = "C:\\Windows\\system32"
	}

	root.RootCmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&CmdFlagOS,
		"os",
		"o",
		runtime.GOOS,
		"OS",
	)
	Cmd.Flags().StringVarP(
		&CmdFlagArch,
		"arch",
		"a",
		runtime.GOARCH,
		"Architecture",
	)
	Cmd.Flags().StringVarP(
		&CmdFlagBinDir,
		"bin-dir",
		"d",
		defaultBinDir,
		"Binary dir",
	)
}
