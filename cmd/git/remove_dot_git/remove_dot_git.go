package remove_dot_git

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	parent_cmd "github.com/sikalabs/slu/cmd/git"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "remove-dot-git",
	Short: "Remove all .git directories in tree",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		// Find all .git directories
		var dotGitDirs []string
		err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() && info.Name() == ".git" {
				dotGitDirs = append(dotGitDirs, path)
				// Don't traverse into .git directories
				return filepath.SkipDir
			}
			return nil
		})
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error walking directory: %v\n", err)
			os.Exit(1)
		}

		if len(dotGitDirs) == 0 {
			fmt.Println("No .git directories found")
			return
		}

		// Print all found .git directories
		fmt.Println("Found .git directories:")
		for _, dir := range dotGitDirs {
			fmt.Printf("  %s\n", dir)
		}

		// Ask for confirmation
		fmt.Printf("\nType 'yes' to remove these directories: ")
		reader := bufio.NewReader(os.Stdin)
		response, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
			os.Exit(1)
		}

		response = strings.TrimSpace(response)
		if response != "yes" {
			fmt.Println("Aborted")
			return
		}

		// Remove all .git directories
		for _, dir := range dotGitDirs {
			err := os.RemoveAll(dir)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error removing %s: %v\n", dir, err)
			} else {
				fmt.Printf("Removed: %s\n", dir)
			}
		}
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}
