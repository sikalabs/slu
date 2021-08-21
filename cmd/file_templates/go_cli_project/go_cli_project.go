package go_cli_project

import (
	"html/template"
	"os"
	"path"
	"path/filepath"

	file_templates_cmd "github.com/sikalabs/slu/cmd/file_templates"

	"github.com/spf13/cobra"
)

var CmdFlagPathPrefix string
var CmdFlagProjectName string
var CmdFlagPackage string

type TemplateVariables struct {
	ProjectName string
	Package     string
}

var Cmd = &cobra.Command{
	Use:   "go-cli-project",
	Short: "Create Golang CLI project",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		for filename, content := range Files {
			_ = content
			fullPath := path.Join(CmdFlagPathPrefix, filename)
			err := os.MkdirAll(filepath.Dir(fullPath), os.ModePerm)
			if err != nil {
				panic(err)
			}
			t, err := template.New(fullPath).Parse(content)
			if err != nil {
				panic(err)
			}
			f, err := os.Create(fullPath)
			if err != nil {
				panic(err)
			}
			defer f.Close()
			t.Execute(f, TemplateVariables{
				ProjectName: CmdFlagProjectName,
				Package:     CmdFlagPackage,
			})
		}
	},
}

func init() {
	file_templates_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVar(
		&CmdFlagPathPrefix,
		"path",
		".",
		"Path prefix",
	)
	Cmd.Flags().StringVarP(
		&CmdFlagProjectName,
		"project-name",
		"n",
		"",
		"Project name {{.ProjectName}}",
	)
	Cmd.MarkFlagRequired("project-name")
	Cmd.Flags().StringVarP(
		&CmdFlagPackage,
		"package",
		"p",
		"",
		"Package {{.Package}}",
	)
	Cmd.MarkFlagRequired("package")
}
