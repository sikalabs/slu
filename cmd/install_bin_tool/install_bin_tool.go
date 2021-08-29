package install_bin_tool

import (
	"bytes"
	"fmt"
	"log"
	"runtime"
	"text/template"

	"github.com/sikalabs/slu/cmd/root"
	"github.com/sikalabs/slu/utils/install_bin_utils"
	"github.com/spf13/cobra"
)

type Tool struct {
	Name        string
	SourcePath  string
	Version     string
	UrlTemplate string
}

var CmdFlagBinDir string
var CmdFlagOS string
var CmdFlagArch string

var Cmd = &cobra.Command{
	Use:   "install-bin-tool",
	Short: "Install preconfigured binary tool like Terraform, Vault, ...",
}

func getUrl(urlTemplate, version string) string {
	tmpl, err := template.New("main").Parse(urlTemplate)
	if err != nil {
		panic(err)
	}
	var out bytes.Buffer
	err = tmpl.Execute(&out, map[string]string{
		"Os":         CmdFlagOS,
		"OsDocker":   dockerOs(CmdFlagOS),
		"Arch":       CmdFlagArch,
		"ArchDocker": dockerArch(CmdFlagArch),
		"Version":    version,
	})
	if err != nil {
		panic(err)
	}
	return out.String()
}

func buildCmd(name, source, url string) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   name,
		Short: "Install " + name + " binary",
		Args:  cobra.NoArgs,
		Run: func(c *cobra.Command, args []string) {
			if source == "" {
				source = name
			}
			install_bin_utils.InstallBin(
				url,
				source,
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
	for _, tool := range Tools {
		Cmd.AddCommand(buildCmd(tool.Name, tool.SourcePath, getUrl(tool.UrlTemplate, tool.Version)))
	}
}

func dockerOs(osName string) string {
	if osName == "darwin" {
		return "mac"
	}
	if osName == "windows" {
		return "win"
	}
	if osName == "linux" {
		return osName
	}
	return ""
}

func dockerArch(arch string) string {
	if arch == "adm64" {
		return "x86_64"
	}
	if arch == "arm64" {
		return "aarch64"
	}
	log.Fatal(fmt.Errorf("unknown arch: %s", arch))
	return ""
}
