package remove_dot_terraform

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	parent_cmd "github.com/sikalabs/slu/cmd/terraform"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "remove-dot-terraform",
	Short: "Remove all .terraform directories in tree",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		var dotTerraformDirs []string
		err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() && info.Name() == ".terraform" {
				dotTerraformDirs = append(dotTerraformDirs, path)
				return filepath.SkipDir
			}
			return nil
		})
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error walking directory: %v\n", err)
			os.Exit(1)
		}

		if len(dotTerraformDirs) == 0 {
			fmt.Println("No .terraform directories found")
			return
		}

		fmt.Println("Found .terraform directories:")
		for _, dir := range dotTerraformDirs {
			fmt.Printf("  %s\n", dir)
		}

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

		for _, dir := range dotTerraformDirs {
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
