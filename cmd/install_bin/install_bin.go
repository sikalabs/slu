package install_bin

import (
	"bytes"
	"fmt"
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
	GetOsFunc      func(string) string
	GetArchFunc    func(string) string
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

func getUrl(
	urlTemplate string,
	version string,
	getOsFunc func(string) string,
	getArchFunc func(string) string,
	defaultOSFunc func() string,
	defaultArchFunc func() string,
) string {
	os_ := getOsFunc(defaultOSFunc())
	arch := getArchFunc(defaultArchFunc())

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
		"Os":         os_,
		"OsDocker":   dockerOs(os_),
		"OsK6":       k6_Os(os_),
		"Arch":       arch,
		"ArchDocker": dockerArch(arch),
		"Version":    version,
	})
	if err != nil {
		panic(err)
	}
	return out.String()
}

func getSourcePath(
	SourcePathTemplate string,
	version string,
	getOsFunc func(string) string,
	getArchFunc func(string) string,
	defaultOSFunc func() string,
	defaultArchFunc func() string,
) string {
	os_ := getOsFunc(defaultOSFunc())
	arch := getArchFunc(defaultArchFunc())

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
		"Os":         os_,
		"OsDocker":   dockerOs(os_),
		"OsK6":       k6_Os(os_),
		"Arch":       arch,
		"ArchDocker": dockerArch(arch),
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
	flagVersionFunc func() string,
	defaultVersionFunc func() string,
	getUrlFunc func(string, string, func(string) string, func(string) string, func() string, func() string) string,
	getSourcePathFunc func(string, string, func(string) string, func(string) string, func() string, func() string) string,
	getOsFunc func(string) string,
	getArchFunc func(string) string,
	defaultOSFunc func() string,
	defaultArchFunc func() string,
	defaultBinDirFunc func() string,
) *cobra.Command {
	var cmd = &cobra.Command{
		Use:     name,
		Short:   "Install " + name + " binary",
		Aliases: aliases,
		Args:    cobra.NoArgs,
		Run: func(c *cobra.Command, args []string) {
			run(
				name,
				aliases,
				sourceTemlate,
				urlTemplate,
				flagVersionFunc,
				defaultVersionFunc,
				getUrlFunc,
				getSourcePathFunc,
				getOsFunc,
				getArchFunc,
				defaultOSFunc,
				defaultArchFunc,
				defaultBinDirFunc,
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
		getOsFunc := func(x string) string { return x }
		if tool.GetOsFunc != nil {
			getOsFunc = tool.GetOsFunc
		}

		getArchFunc := func(x string) string { return x }
		if tool.GetArchFunc != nil {
			getArchFunc = tool.GetArchFunc
		}
		Cmd.AddCommand(buildCmd(
			tool.Name,
			tool.Aliases,
			tool.SourcePath,
			tool.UrlTemplate,
			func() string { return FlagVersion },
			tool.GetVersionFunc,
			getUrl,
			getSourcePath,
			getOsFunc,
			getArchFunc,
			func() string { return CmdFlagOS },
			func() string { return CmdFlagArch },
			func() string { return CmdFlagBinDir },
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
	return ""
}

func k6_Os(osName string) string {
	if osName == "darwin" {
		return "macos"
	}
	return osName
}

func run(
	name string,
	aliases []string,
	sourceTemlate string,
	urlTemplate string,
	flagVersionFunc func() string,
	defaultVersionFunc func() string,
	getUrlFunc func(string, string, func(string) string, func(string) string, func() string, func() string) string,
	getSourcePathFunc func(string, string, func(string) string, func(string) string, func() string, func() string) string,
	getOsFunc func(string) string,
	getArchFunc func(string) string,
	defaultOSFunc func() string,
	defaultArchFunc func() string,
	defaultBinDirFunc func() string,
) {
	if sourceTemlate == "" {
		sourceTemlate = name
	}
	version := ""
	if flagVersionFunc() != "latest" {
		version = flagVersionFunc()
	} else {
		version = defaultVersionFunc()
	}
	url := getUrlFunc(
		urlTemplate,
		version,
		getOsFunc,
		getArchFunc,
		defaultOSFunc,
		defaultArchFunc,
	)
	if FlagVerbose {
		fmt.Println(url)
	}
	source := getSourcePathFunc(
		sourceTemlate,
		version,
		getOsFunc,
		getArchFunc,
		defaultOSFunc,
		defaultArchFunc,
	)
	install_bin_utils.InstallBin(
		url,
		source,
		defaultBinDirFunc(),
		name,
		defaultOSFunc() == "windows",
	)
}
