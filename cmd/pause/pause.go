package pause

import (
	"github.com/sikalabs/pause/pkg/pause"
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:  "pause",
	Args: cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		pause.Pause()
	},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
