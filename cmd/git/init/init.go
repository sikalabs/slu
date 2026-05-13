package init

import (
	"fmt"
	"os"
	"path/filepath"

	parent_cmd "github.com/sikalabs/slu/cmd/git"
	"github.com/sikalabs/slu/internal/error_utils"
	"github.com/sikalabs/slu/utils/exec_utils"
	"github.com/spf13/cobra"
)

var FlagNameOverride string

var Cmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize Git repository with README",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		// Get current directory name
		cwd, err := os.Getwd()
		error_utils.HandleError(err, "Error getting current directory")
		name := filepath.Base(cwd)
		if FlagNameOverride != "" {
			name = FlagNameOverride
		}

		// Initialize git repository
		err = exec_utils.ExecOut("git", "init")
		error_utils.HandleError(err, "Error initializing git repository")

		// Create README.md with directory name as header
		readmeContent := fmt.Sprintf("# %s\n", name)
		err = os.WriteFile("README.md", []byte(readmeContent), 0644)
		error_utils.HandleError(err, "Error creating README.md")

		// Add README.md to staging
		err = exec_utils.ExecOut("git", "add", "README.md")
		error_utils.HandleError(err, "Error adding README.md")

		// Commit with initial message
		commitMsg := fmt.Sprintf("init: Initial commit, %s", name)
		err = exec_utils.ExecOut("git", "commit", "-m", commitMsg)
		error_utils.HandleError(err, "Error creating initial commit")
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)

	Cmd.Flags().StringVar(
		&FlagNameOverride,
		"name-override",
		"",
		"Override the name used in README and initial commit",
	)
}
