package ip

import (
	"os"
	"os/exec"
	"strconv"

	"github.com/sikalabs/slu/cmd/root"
	"github.com/sikalabs/slu/utils/time_utils"
	"github.com/spf13/cobra"
)

var FlagJson bool

var Cmd = &cobra.Command{
	Use:   "ddev",
	Short: "Run sikalabs/dev in Docker",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		cmd := exec.Command(
			"docker", "run",
			"--name", "dev-"+strconv.Itoa(time_utils.Unix()),
			"--rm", "-ti",
			"sikalabs/dev",
			"bash",
		)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		cmd.Run()
	},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
