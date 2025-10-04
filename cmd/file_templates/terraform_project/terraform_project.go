package terraform_project

import (
	"embed"
	"io"
	"io/fs"
	"os"
	"path"

	file_templates_cmd "github.com/sikalabs/slu/cmd/file_templates"
	"github.com/sikalabs/slu/internal/error_utils"

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
	error_utils.HandleError(err, "Failed to get sub filesystem")
	fs.WalkDir(secFS, ".", func(filePath string, d fs.DirEntry, err error) error {
		error_utils.HandleError(err, "Failed to walk directory")
		if d.Type().IsDir() {
			os.MkdirAll(path.Join(CmdFlagPathPrefix, filePath), 0755)
		} else {
			file, err := secFS.Open(filePath)
			error_utils.HandleError(err, "Failed to open file")
			fileContent, err := io.ReadAll(file)
			error_utils.HandleError(err, "Failed to read file")
			if contains(executables, filePath) {
				err = os.WriteFile(path.Join(CmdFlagPathPrefix, filePath), fileContent, 0755)
			} else {
				err = os.WriteFile(path.Join(CmdFlagPathPrefix, filePath), fileContent, 0644)
			}
			error_utils.HandleError(err, "Failed to write file")
		}
		return nil
	})
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
