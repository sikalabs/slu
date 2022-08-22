package terraform_project

import (
	"embed"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path"

	file_templates_cmd "github.com/sikalabs/slu/cmd/file_templates"

	"github.com/spf13/cobra"
)

var CmdFlagPathPrefix string

type TemplateVariables struct{}

//go:embed _data/*
//go:embed _data/**/*
var srcDataFS embed.FS

var executables = []string{
	".git-hooks/pre-commit",
}

var Cmd = &cobra.Command{
	Use:     "terraform-project",
	Short:   "Create Terraform Project",
	Aliases: []string{"tfp"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		do()
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
}

func do() {
	var err error
	secFS, err := fs.Sub(srcDataFS, "_data")
	handleError(err)
	fs.WalkDir(secFS, ".", func(filePath string, d fs.DirEntry, err error) error {
		handleError(err)
		if d.Type().IsDir() {
			os.MkdirAll(path.Join(CmdFlagPathPrefix, filePath), 0755)
		} else {
			file, err := secFS.Open(filePath)
			handleError(err)
			fileContent, err := ioutil.ReadAll(file)
			handleError(err)
			if contains(executables, filePath) {
				err = ioutil.WriteFile(path.Join(CmdFlagPathPrefix, filePath), fileContent, 0755)
			} else {
				err = ioutil.WriteFile(path.Join(CmdFlagPathPrefix, filePath), fileContent, 0644)
			}
			handleError(err)
		}
		return nil
	})
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
