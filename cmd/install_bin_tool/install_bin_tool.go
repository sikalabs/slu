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

func installHashicorpBin(tool, version string) {
	url := "https://releases.hashicorp.com/" + tool + "/" +
		version + "/terraform_" + version + "_" + CmdFlagOS + "_" + CmdFlagArch + ".zip"
	install_bin_utils.InstallBin(
		url,
		"terraform",
		CmdFlagBinDir,
		"terraform",
	)
}

var InstallTerraformCmd = &cobra.Command{
	Use:   "terraform",
	Short: "Install terraform binary",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		installHashicorpBin("terraform", "1.0.5")
	},
}

var VaultTerraformCmd = &cobra.Command{
	Use:   "vault",
	Short: "Install vault binary",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		installHashicorpBin("vault", "1.8.2")
	},
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
	Cmd.AddCommand(InstallTerraformCmd)
	Cmd.AddCommand(VaultTerraformCmd)
}
