package new_repo

import (
	parent_cmd "github.com/sikalabs/slu/cmd/scripts"
	"github.com/sikalabs/slu/internal/error_utils"
	"github.com/sikalabs/slu/utils/exec_utils"
	"github.com/spf13/cobra"
)

var FlagWithGo bool
var FlagWithPython bool
var FlagTerraform bool
var FlagKubernetes bool
var FlagNodeJS bool
var FlagNextJs bool
var FlagHelm bool

var Cmd = &cobra.Command{
	Use:     "new-repo",
	Short:   "Initialize Git repository with README, .editorconfig and .gitignore",
	Aliases: []string{"nr"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		// Step 1: Initialize git repository
		err := exec_utils.ExecOut("slu", "git", "init")
		error_utils.HandleError(err, "Error running slu git init")

		// Step 2: Create .editorconfig
		editorconfigArgs := []string{"ft", "ec"}
		if FlagWithGo {
			editorconfigArgs = append(editorconfigArgs, "--go")
		}
		if FlagWithPython {
			editorconfigArgs = append(editorconfigArgs, "--python")
		}
		err = exec_utils.ExecOut("slu", editorconfigArgs...)
		error_utils.HandleError(err, "Error running slu ft ec")

		// Step 3: Create .gitignore
		gitignoreArgs := []string{"ft", "gi"}
		if FlagTerraform {
			gitignoreArgs = append(gitignoreArgs, "--terraform")
		}
		if FlagKubernetes {
			gitignoreArgs = append(gitignoreArgs, "--kubernetes")
		}
		if FlagNodeJS {
			gitignoreArgs = append(gitignoreArgs, "--node")
		}
		if FlagNextJs {
			gitignoreArgs = append(gitignoreArgs, "--nextjs")
		}
		if FlagHelm {
			gitignoreArgs = append(gitignoreArgs, "--helm")
		}
		err = exec_utils.ExecOut("slu", gitignoreArgs...)
		error_utils.HandleError(err, "Error running slu ft gi")

		// Step 4: Commit .editorconfig and .gitignore
		err = exec_utils.ExecOut("slu", "git", "commit", "editorconfig-and-gitignore")
		error_utils.HandleError(err, "Error running slu git commit editorconfig-and-gitignore")
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)

	// Flags from editorconfig command
	Cmd.Flags().BoolVar(
		&FlagWithGo,
		"go",
		false,
		"Add Go section to editorconfig",
	)
	Cmd.Flags().BoolVar(
		&FlagWithPython,
		"python",
		false,
		"Add Python section to editorconfig",
	)

	// Flags from gitignore command
	Cmd.Flags().BoolVar(
		&FlagTerraform,
		"terraform",
		false,
		"Add Terraform part to gitignore",
	)
	Cmd.Flags().BoolVar(
		&FlagKubernetes,
		"kubernetes",
		false,
		"Add Kubernetes part to gitignore",
	)
	Cmd.Flags().BoolVar(
		&FlagNodeJS,
		"node",
		false,
		"Add NodeJS part to gitignore",
	)
	Cmd.Flags().BoolVar(
		&FlagNextJs,
		"nextjs",
		false,
		"Add Next.js part to gitignore",
	)
	Cmd.Flags().BoolVar(
		&FlagHelm,
		"helm",
		false,
		"Add Helm part to gitignore",
	)
}
