package status

import (
	"fmt"

	lock_cmd "github.com/sikalabs/slu/cmd/k8s/lock"
	"github.com/sikalabs/slu/utils/k8s_lock_utils"

	"github.com/spf13/cobra"
)

var CmdFlagName string
var CmdFlagNamespace string

var Cmd = &cobra.Command{
	Use:   "status",
	Short: "Get status of lock in Kubernetes",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		isLocked := k8s_lock_utils.CheckLock(CmdFlagName, CmdFlagNamespace)
		fmt.Println("is-locked", isLocked)
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
