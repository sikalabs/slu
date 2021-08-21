package go_cli_project

var Files = map[string]string{
	// README.md
	"README.md": `# {{.ProjectName}}
`,
	".gitignore": `# Mac
.DS_Store

# Editor
.vscode
.idea

# Generic
*.log
*.backup

# Go
{{.ProjectName}}
*.exe
/dist/**
cobra-docs
`,
	// .editorconfig
	".editorconfig": `root = true
[*]
indent_style = space
indent_size = 2
charset = utf-8
trim_trailing_whitespace = true
insert_final_newline = true
end_of_line = lf
max_line_length = off
[*.go]
indent_style = tab
[Makefile]
indent_style = tab
`,
	// go.mod
	"go.mod": `module {{.Package}}

go 1.16

require github.com/spf13/cobra v1.2.1
`,
	// version/version.go
	"version/version.go": `package version

var Version string = "v0.1.0-dev"
`,
	// cmd/cmd.go
	"cmd/cmd.go": `package cmd

import (
	"{{.Package}}/cmd/root"
	_ "{{.Package}}/cmd/version"
	"github.com/spf13/cobra"
)

func Execute() {
	cobra.CheckErr(root.Cmd.Execute())
}
`,
	// cmd/root/root.go
	"cmd/root/root.go": `
package root

import (
	"{{.Package}}/version"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "{{.ProjectName}}",
	Short: "{{.ProjectName}}, " + version.Version,
}
`,
	// cmd/version/version.go
	"cmd/version/version.go": `package version

import (
	"fmt"

	"{{.Package}}/cmd/root"
	"{{.Package}}/version"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "version",
	Short:   "Prints version",
	Aliases: []string{"v"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		fmt.Printf("%s\n", version.Version)
	},
}

func init() {
	root.Cmd.AddCommand(Cmd)
}
`,
	// main.go
	"main.go": `package main

import (
	"{{.Package}}/cmd"
)

func main() {
	cmd.Execute()
}
`,
}
