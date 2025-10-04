package version

import (
	"strings"

	parent_cmd "github.com/sikalabs/slu/cmd/check"
	"github.com/sikalabs/slu/internal/error_utils"
	"github.com/sikalabs/slu/utils/exec_utils"

	"github.com/spf13/cobra"
)

var FlagPrefix string

var Cmd = &cobra.Command{
	Use:   "kubernetes_context",
	Short: "Kubernetes context must have prefix",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		currentContextRaw, err := exec_utils.ExecStr("kubectl", "config", "current-context")
		error_utils.HandleError(err, "Failed to get current kubernetes context")
		currentContex := strings.ReplaceAll(currentContextRaw, "\n", "")
		ok := strings.HasPrefix(currentContex, FlagPrefix)
		error_utils.HandleNotOK(ok, "Kubernetes context ("+currentContex+") must have prefix "+FlagPrefix)
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagPrefix,
		"prefix",
		"p",
		"",
		"prefix",
	)
	Cmd.MarkFlagRequired("prefix")
}
