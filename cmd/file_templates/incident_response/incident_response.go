package incident_response

import (
	"time"

	file_templates_cmd "github.com/sikalabs/slu/cmd/file_templates"

	"github.com/spf13/cobra"
)

var FlagTitle string
var FlagDate string
var FlagPathPrefix string

var Cmd = &cobra.Command{
	Use:     "incident-response",
	Short:   "Create basic gitignore",
	Aliases: []string{"ir"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		if FlagDate == "today" {
			FlagDate = time.Now().Format("2006-01-02")
		}
		CreateIncidentResponseFile(FlagPathPrefix, FlagDate, FlagTitle)
	},
}

func init() {
	file_templates_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagTitle,
		"title",
		"t",
		"",
		"Title of incident",
	)
	Cmd.MarkFlagRequired("title")
	Cmd.Flags().StringVarP(
		&FlagDate,
		"date",
		"d",
		"today",
		"Incident date in format yyyy-mm-dd",
	)
	Cmd.Flags().StringVarP(
		&FlagPathPrefix,
		"path-prefix",
		"p",
		".",
		"Path prefix",
	)
}
