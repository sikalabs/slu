package google_drive

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "google-drive",
	Short:   "Google Drive Utils",
	Aliases: []string{"gdrive", "gdr"},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
