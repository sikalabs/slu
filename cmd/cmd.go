package cmd

import (
	_ "github.com/sikalabs/slu/cmd/expand"
	_ "github.com/sikalabs/slu/cmd/expand/file"
	_ "github.com/sikalabs/slu/cmd/expand/string"
	_ "github.com/sikalabs/slu/cmd/file_templates"
	_ "github.com/sikalabs/slu/cmd/file_templates/editorconfig"
	_ "github.com/sikalabs/slu/cmd/file_templates/gitignore"
	_ "github.com/sikalabs/slu/cmd/generate_docs"
	_ "github.com/sikalabs/slu/cmd/go_code"
	_ "github.com/sikalabs/slu/cmd/go_code/version_bump"
	_ "github.com/sikalabs/slu/cmd/k8s"
	_ "github.com/sikalabs/slu/cmd/k8s/get"
	_ "github.com/sikalabs/slu/cmd/k8s/get/secret"
	_ "github.com/sikalabs/slu/cmd/k8s/token"
	_ "github.com/sikalabs/slu/cmd/mysql"
	_ "github.com/sikalabs/slu/cmd/mysql/create"
	_ "github.com/sikalabs/slu/cmd/mysql/drop"
	_ "github.com/sikalabs/slu/cmd/mysql/list"
	_ "github.com/sikalabs/slu/cmd/postgres"
	_ "github.com/sikalabs/slu/cmd/postgres/create"
	_ "github.com/sikalabs/slu/cmd/postgres/drop"
	_ "github.com/sikalabs/slu/cmd/postgres/list"
	"github.com/sikalabs/slu/cmd/root"
	_ "github.com/sikalabs/slu/cmd/version"
	"github.com/spf13/cobra"
)

func Execute() {
	cobra.CheckErr(root.RootCmd.Execute())
}
