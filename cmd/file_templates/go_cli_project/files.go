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
	// .goreleaser.yml
	".goreleaser.yml": `project_name: {{.ProjectName}}

before:
  hooks:
    - rm -rf ./dist
    - go mod tidy
builds:
  -
    env:
      - CGO_ENABLED=0
    mod_timestamp: "{{"{{"}} .CommitTimestamp {{"}}"}}"
    flags:
      - -trimpath
    ldflags:
      - -s
      - -w
      - -X {{.Package}}/version.Version=v{{"{{"}}.Version{{"}}"}}
    goos:
      - windows
      - linux
      - darwin
    goarch:
      - amd64
      - "386"
      - arm
      - arm64
    goarm:
      - 6
      - 7
    ignore:
      - goos: darwin
        goarch: "386"
      - goos: windows
        goarch: "arm"
      - goos: windows
        goarch: "arm64"
      - goos: linux
        goarch: arm
        goarm: 6
    binary: {{.ProjectName}}

archives:
  - format: tar.gz
    name_template: "{{"{{"}} .ProjectName {{"}}"}}_v{{"{{"}} .Version {{"}}"}}_{{"{{"}} .Os {{"}}"}}_{{"{{"}} .Arch {{"}}"}}"

release:
  prerelease: auto

checksum:
  name_template: "{{"{{"}} .ProjectName {{"}}"}}_checksums.txt"
  algorithm: sha256

brews:
  -
    name: {{.ProjectName}}
    conflicts:
      - {{.ProjectName}}-edge
    tap:
      owner: {{.BrewOrganization}}
      name: {{.BrewRepo}}
    skip_upload: auto
    homepage: https://{{.Package}}
    url_template: "https://{{.Package}}/releases/download/{{"{{"}} .Tag {{"}}"}}/{{"{{"}} .ArtifactName {{"}}"}}"
    folder: Formula
    caveats: "How to use this binary: https://{{.Package}}"
    description: "{{.ProjectName}}"
    install: |
      bin.install "{{.ProjectName}}"
    test: |
      system "#{bin}/{{.ProjectName}} version"
  -
    name: {{.ProjectName}}-edge
    conflicts:
      - {{.ProjectName}}
    tap:
      owner: {{.BrewOrganization}}
      name: {{.BrewRepo}}
    skip_upload: false
    homepage: https://{{.Package}}
    url_template: "https://{{.Package}}/releases/download/{{"{{"}} .Tag {{"}}"}}/{{"{{"}} .ArtifactName {{"}}"}}"
    folder: Formula
    caveats: "How to use this binary: https://{{.Package}}"
    description: "{{.ProjectName}}"
    install: |
      bin.install "{{.ProjectName}}"
    test: |
      system "#{bin}/{{.ProjectName}} version"

dockers:
    -
      goos: linux
      goarch: amd64
      image_templates:
        - "{{.DockerRegistry}}/{{.ProjectName}}:{{"{{"}} .Tag {{"}}"}}"
      dockerfile: Dockerfile
      ids:
        - {{.ProjectName}}
      build_flag_templates:
        - "--platform=linux/amd64"
        - "--label=org.opencontainers.image.created={{"{{"}}.Date{{"}}"}}"
        - "--label=org.opencontainers.image.title={{"{{"}}.ProjectName{{"}}"}}"
        - "--label=org.opencontainers.image.revision={{"{{"}}.FullCommit{{"}}"}}"
        - "--label=org.opencontainers.image.version={{"{{"}}.Version{{"}}"}}"
        - "--label=org.label-schema.schema-version=1.0"
        - "--label=org.label-schema.version={{"{{"}}.Version{{"}}"}}"
        - "--label=org.label-schema.name={{"{{"}}.ProjectName{{"}}"}}"
        - "--label=com.github.actions.name={{"{{"}}.ProjectName{{"}}"}}"
        - "--label=repository=https://{{.Package}}"
        {{ if .Maintainer }}- "--label=maintainer={{.Maintainer}}"{{ end }}
    - goos: linux
      goarch: arm64
      image_templates:
        - "{{.DockerRegistry}}/{{.ProjectName}}:{{"{{"}} .Tag {{"}}"}}-arm64v8"
      dockerfile: Dockerfile.arm64v8
      ids:
        - {{.ProjectName}}
      build_flag_templates:
        - "--platform=linux/arm64"
        - "--label=org.opencontainers.image.created={{"{{"}}.Date{{"}}"}}"
        - "--label=org.opencontainers.image.title={{"{{"}}.ProjectName{{"}}"}}"
        - "--label=org.opencontainers.image.revision={{"{{"}}.FullCommit{{"}}"}}"
        - "--label=org.opencontainers.image.version={{"{{"}}.Version{{"}}"}}"
        - "--label=org.label-schema.schema-version=1.0"
        - "--label=org.label-schema.version={{"{{"}}.Version{{"}}"}}"
        - "--label=org.label-schema.name={{"{{"}}.ProjectName{{"}}"}}"
        - "--label=com.github.actions.name={{"{{"}}.ProjectName{{"}}"}}"
        - "--label=repository=https://{{.Package}}"
        {{ if .Maintainer }}- "--label=maintainer={{.Maintainer}}"{{ end }}
`,
	// Dockerfile
	"Dockerfile": `FROM debian:10-slim
COPY {{.ProjectName}} /usr/local/bin/
`,
	// Dockerfile.arm64v8
	"Dockerfile.arm64v8": `FROM arm64v8/debian:10-slim
COPY {{.ProjectName}} /usr/local/bin/
`,
}
