package install_bin_tool

import (
	"runtime"

	"github.com/sikalabs/slu/cmd/root"
	"github.com/sikalabs/slu/utils/install_bin_utils"
	"github.com/spf13/cobra"
)

var CmdFlagBinDir string
var CmdFlagOS string
var CmdFlagArch string

var Cmd = &cobra.Command{
	Use:   "install-bin-tool",
	Short: "Install preconfigured binary tool like Terraform, Vault, ...",
}

func hashicorpUrl(tool, version string) string {
	return "https://releases.hashicorp.com/" + tool + "/" +
		version + "/terraform_" + version + "_" + CmdFlagOS + "_" + CmdFlagArch + ".zip"
}

func buildCmd(name string, urlFunc func() string) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   name,
		Short: "Install " + name + " binary",
		Args:  cobra.NoArgs,
		Run: func(c *cobra.Command, args []string) {
			install_bin_utils.InstallBin(
				urlFunc(),
				name,
				CmdFlagBinDir,
				name,
			)
		},
	}
	return cmd
}

func init() {
	root.RootCmd.AddCommand(Cmd)
	Cmd.PersistentFlags().StringVarP(
		&CmdFlagBinDir,
		"bin-dir",
		"d",
		"/usr/local/bin",
		"Binary dir",
	)
	Cmd.PersistentFlags().StringVarP(
		&CmdFlagOS,
		"os",
		"o",
		runtime.GOOS,
		"OS",
	)
	Cmd.PersistentFlags().StringVarP(
		&CmdFlagArch,
		"arch",
		"a",
		runtime.GOARCH,
		"Architecture",
	)
	Cmd.AddCommand(buildCmd(
		"terraform",
		func() string { return hashicorpUrl("terraform", "1.0.5") }),
	)
	Cmd.AddCommand(buildCmd(
		"vault",
		func() string { return hashicorpUrl("vault", "1.8.2") }),
	)
}
