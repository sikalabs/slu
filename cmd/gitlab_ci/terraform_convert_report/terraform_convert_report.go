package terraform_convert_report

import (
	parent_cmd "github.com/sikalabs/slu/cmd/gitlab_ci"
	"github.com/sikalabs/slu/utils/exec_utils"

	"github.com/spf13/cobra"
)

const CONVERT_JQ_ARG = `([.resource_changes[]?.change.actions?]|flatten)|{"create":(map(select(.=="create"))|length),"update":(map(select(.=="update"))|length),"delete":(map(select(.=="delete"))|length)}`

var Cmd = &cobra.Command{
	Use:     "terraform-convert-report [<plan.json>]",
	Short:   "Covert Report from Terraform (plan.json)",
	Aliases: []string{"tcr"},
	Args:    cobra.ExactArgs(1),
	Run: func(c *cobra.Command, args []string) {
		path := args[0]
		exec_utils.ExecOut("jq", CONVERT_JQ_ARG, path)
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}
