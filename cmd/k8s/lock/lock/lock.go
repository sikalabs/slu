package lock

import (
	"os"

	lock_cmd "github.com/sikalabs/slu/cmd/k8s/lock"
	"github.com/sikalabs/slu/utils/k8s_lock_utils"

	"github.com/spf13/cobra"
)

var CmdFlagName string
var CmdFlagNamespace string

var Cmd = &cobra.Command{
	Use:   "lock",
	Short: "Create lock in Kubernetes",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		locked, err := k8s_lock_utils.Lock(
			CmdFlagName, CmdFlagNamespace,
		)
		if err != nil {
			panic(err)
		}
		if locked {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	},
}

func init() {
	lock_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&CmdFlagName,
		"lock-name",
		"l",
		"",
		"ConfigMap Name",
	)
	Cmd.MarkFlagRequired("lock-name")
	Cmd.Flags().StringVarP(
		&CmdFlagNamespace,
		"namespace",
		"n",
		"",
		"Kubernetes Namespace",
	)
}
