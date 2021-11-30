package get

import (
	wait_for_cmd "github.com/sikalabs/slu/cmd/wait_for"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "k8s",
	Short:   "Wait for k8s (pods, jobs, ...)",
	Aliases: []string{"k"},
}

func init() {
	wait_for_cmd.Cmd.AddCommand(Cmd)
}
