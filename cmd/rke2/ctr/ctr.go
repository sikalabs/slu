package ctr

import (
	"log"

	parent_cmd "github.com/sikalabs/slu/cmd/rke2"
	"github.com/sikalabs/slu/utils/exec_utils"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "ctr",
	Short: "Run `ctr` with connection to RKE2 containerd",
	Run: func(c *cobra.Command, args []string) {
		ctrBin := "/var/lib/rancher/rke2/bin/ctr"
		ctrArgs := []string{
			"--address", "/run/k3s/containerd/containerd.sock",
			"--namespace", "k8s.io",
		}
		ctrArgs = append(ctrArgs, args...)
		err := exec_utils.ExecOut(ctrBin, ctrArgs...)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}
