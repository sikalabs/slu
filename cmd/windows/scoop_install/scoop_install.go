package scoop_install

import (
	parent_cmd "github.com/sikalabs/slu/cmd/windows"
	"github.com/sikalabs/slu/utils/exec_utils"
	"github.com/spf13/cobra"
)

type Packages struct {
	Name    string
	Aliases []string
	Url     string
}

var Cmd = &cobra.Command{
	Use:   "scoop-install",
	Short: "Install preconfigured binary tool like Terraform, Vault, ...",
}

func buildCmd(
	name string,
	aliases []string,
	url string,
) *cobra.Command {
	var cmd = &cobra.Command{
		Use:     name,
		Short:   "Install " + name + " binary",
		Aliases: aliases,
		Args:    cobra.NoArgs,
		Run: func(c *cobra.Command, args []string) {
			exec_utils.ExecOut("scoop", "install", url)
		},
	}
	return cmd
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
	for _, tool := range Tools {
		Cmd.AddCommand(buildCmd(
			tool.Name,
			tool.Aliases,
			tool.Url,
		))
	}
}
