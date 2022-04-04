package large_desktop_files

import (
	"fmt"
	"os"

	parentcmd "github.com/sikalabs/slu/cmd/ondrejsika"
	"github.com/sikalabs/slu/utils/du_utils"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "large-desktop-files",
	Short: "Find large desktop files",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		dir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		du_utils.RunDiskUsage(true, "1G", dir, 0, true)
	},
}

func init() {
	parentcmd.Cmd.AddCommand(Cmd)
}
