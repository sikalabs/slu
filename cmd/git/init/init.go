package init

import (
	"fmt"
	"os"
	"path/filepath"

	parent_cmd "github.com/sikalabs/slu/cmd/git"
	"github.com/sikalabs/slu/utils/exec_utils"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize Git repository with README",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		// Get current directory name
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Println("Error getting current directory:", err)
			os.Exit(1)
		}
		dirName := filepath.Base(cwd)

		// Initialize git repository
		err = exec_utils.ExecOut("git", "init")
		if err != nil {
			fmt.Println("Error initializing git repository:", err)
			os.Exit(1)
		}

		// Create README.md with directory name as header
		readmeContent := fmt.Sprintf("# %s\n", dirName)
		err = os.WriteFile("README.md", []byte(readmeContent), 0644)
		if err != nil {
			fmt.Println("Error creating README.md:", err)
			os.Exit(1)
		}

		// Add README.md to staging
		err = exec_utils.ExecOut("git", "add", "README.md")
		if err != nil {
			fmt.Println("Error adding README.md:", err)
			os.Exit(1)
		}

		// Commit with initial message
		commitMsg := fmt.Sprintf("init: Initial commit, %s", dirName)
		err = exec_utils.ExecOut("git", "commit", "-m", commitMsg)
		if err != nil {
			fmt.Println("Error creating initial commit:", err)
			os.Exit(1)
		}
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}