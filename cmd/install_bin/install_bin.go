package install_bin

import (
	"bytes"
	"fmt"
	"log"
	"runtime"
	"strings"
	"text/template"

	"github.com/sikalabs/slu/cmd/root"
	"github.com/sikalabs/slu/utils/install_bin_utils"
	"github.com/spf13/cobra"
)

type Tool struct {
	Name           string
	Aliases        []string
	SourcePath     string
	GetVersionFunc func() string
	UrlTemplate    string
}

var CmdFlagBinDir string
var CmdFlagOS string
var CmdFlagArch string
var FlagVersion string
var FlagVerbose bool

var Cmd = &cobra.Command{
	Use:   "install-bin",
	Short: "Install preconfigured binary tool like Terraform, Vault, ...",
	Aliases: []string{
		"ib",
		// Deprecated aliases
		"install-bin-tool",
		"ibt",
	},
}

func getUrl(urlTemplate, version string) string {
	funcMap := template.FuncMap{
		"capitalize": strings.Title,
		"removev": func(s string) string {
			return strings.ReplaceAll(s, "v", "")
		},
	}
	tmpl, err := template.New("main").Funcs(funcMap).Parse(urlTemplate)
	if err != nil {
		panic(err)
	}
	var out bytes.Buffer
	err = tmpl.Execute(&out, map[string]string{
		"Os":         CmdFlagOS,
		"OsDocker":   dockerOs(CmdFlagOS),
		"OsK6":       k6_Os(CmdFlagOS),
		"Arch":       CmdFlagArch,
		"ArchDocker": dockerArch(CmdFlagArch),
		"ArchK9s":    k9sArch(CmdFlagArch),
		"Version":    version,
	})
	if err != nil {
		panic(err)
	}
	return out.String()
}

func getSourcePath(SourcePathTemplate, version string) string {
	funcMap := template.FuncMap{
		"capitalize": strings.Title,
		"removev": func(s string) string {
			return strings.ReplaceAll(s, "v", "")
		},
	}
	tmpl, err := template.New("source-path").Funcs(funcMap).Parse(SourcePathTemplate)
	if err != nil {
		panic(err)
	}
	var out bytes.Buffer
	err = tmpl.Execute(&out, map[string]string{
		"Os":         CmdFlagOS,
		"OsDocker":   dockerOs(CmdFlagOS),
		"OsK6":       k6_Os(CmdFlagOS),
		"Arch":       CmdFlagArch,
		"ArchDocker": dockerArch(CmdFlagArch),
		"ArchK9s":    k9sArch(CmdFlagArch),
		"Version":    version,
	})
	if err != nil {
		panic(err)
	}
	return out.String()
}

func buildCmd(
	name string,
	aliases []string,
	sourceTemlate string,
	urlTemplate string,
	defaultVersionFunc func() string,
	getUrlFunc func(string, string) string,
	getSourcePathFunc func(string, string) string,
) *cobra.Command {
	var cmd = &cobra.Command{
		Use:     name,
		Short:   "Install " + name + " binary",
		Aliases: aliases,
		Args:    cobra.NoArgs,
		Run: func(c *cobra.Command, args []string) {
			if sourceTemlate == "" {
				sourceTemlate = name
			}
			version := defaultVersionFunc()
			if FlagVersion != "latest" {
				version = FlagVersion
			}
			url := getUrlFunc(urlTemplate, version)
			if FlagVerbose {
				fmt.Println(url)
			}
			source := getSourcePathFunc(sourceTemlate, version)
			install_bin_utils.InstallBin(
				url,
				source,
				CmdFlagBinDir,
				name,
				CmdFlagOS == "windows",
			)
		},
	}
	return cmd
}

func init() {
	defaultBinDir := "/usr/local/bin"
	if runtime.GOOS == "windows" {
		defaultBinDir = "C:\\Windows\\system32"
	}

	root.RootCmd.AddCommand(Cmd)
	Cmd.PersistentFlags().StringVarP(
		&CmdFlagBinDir,
		"bin-dir",
		"d",
		defaultBinDir,
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
	Cmd.PersistentFlags().StringVarP(
		&FlagVersion,
		"version",
		"v",
		"latest",
		"Version",
	)
	Cmd.PersistentFlags().BoolVar(
		&FlagVerbose,
		"verbose",
		false,
		"Verbose output",
	)
	for _, tool := range Tools {
		Cmd.AddCommand(buildCmd(
			tool.Name,
			tool.Aliases,
			tool.SourcePath,
			tool.UrlTemplate,
			tool.GetVersionFunc,
			getUrl,
			getSourcePath,
		))
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
	if arch == "amd64" {
		return "x86_64"
	}
	if arch == "arm64" {
		return "aarch64"
	}
	log.Fatal(fmt.Errorf("unknown arch: %s", arch))
	return ""
}

func k9sArch(arch string) string {
	if arch == "amd64" {
		return "x86_64"
	}
	return arch
}

func k6_Os(osName string) string {
	if osName == "darwin" {
		return "macos"
	}
	return osName
}
