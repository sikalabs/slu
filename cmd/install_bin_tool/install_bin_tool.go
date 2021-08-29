package install_bin_tool

import (
	"runtime"

	"github.com/sikalabs/slu/cmd/root"
	"github.com/sikalabs/slu/utils/install_bin_utils"
	"github.com/spf13/cobra"
)

var CmdFlagBinDir string

var Cmd = &cobra.Command{
	Use:   "install-bin-tool",
	Short: "Install preconfigured binary tool like Terraform, Vault, ...",
}

var InstallTerraformCmd = &cobra.Command{
	Use:   "terraform",
	Short: "Install terraform binary",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		version := "1.0.5"
		arch := runtime.GOOS
		url := "https://releases.hashicorp.com/terraform/" +
			version + "/terraform_" + version + "_" + arch + "_amd64.zip"
		install_bin_utils.InstallBin(
			url,
			"terraform",
			CmdFlagBinDir,
			"terraform",
		)
	},
}

var VaultTerraformCmd = &cobra.Command{
	Use:   "vault",
	Short: "Install vault binary",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		version := "1.8.2"
		arch := runtime.GOOS
		url := "https://releases.hashicorp.com/vault/" +
			version + "/vault_" + version + "_" + arch + "_amd64.zip"
		install_bin_utils.InstallBin(
			url,
			"vault",
			CmdFlagBinDir,
			"vault",
		)
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
	Cmd.AddCommand(InstallTerraformCmd)
	Cmd.AddCommand(VaultTerraformCmd)
}
