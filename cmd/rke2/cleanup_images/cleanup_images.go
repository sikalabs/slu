package cleanup_images

import (
	parent_cmd "github.com/sikalabs/slu/cmd/rke2"
	"github.com/sikalabs/slu/utils/exec_utils"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "cleanup-images",
	Short: "Cleanup images from RKE2 using crictl",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		crictlBin := "/var/lib/rancher/rke2/bin/crictl"
		crictlArgs := []string{
			"--runtime-endpoint", "unix:///run/k3s/containerd/containerd.sock",
			"rmi", "-a",
		}
		exec_utils.ExecOut(crictlBin, crictlArgs...)
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}
