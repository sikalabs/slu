package terraform_lock

import (
	commit_cmd "github.com/sikalabs/slu/cmd/git/commit"
	"github.com/sikalabs/slu/utils/exec_utils"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "terraform-lock",
	Short:   "Commit .terraform.lock.hcl",
	Aliases: []string{"tfl"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		exec_utils.ExecOut("git", "add", ".terraform.lock.hcl")
		exec_utils.ExecOut(
			"git", "commit", "-n", "-m",
			"[auto] chore(terraform.lock.hcl): Update Terraform lock",
			".terraform.lock.hcl")
	},
}

func init() {
	commit_cmd.Cmd.AddCommand(Cmd)
}
