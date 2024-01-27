package spaces_and_commas

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	parent_cmd "github.com/sikalabs/slu/cmd/scripts/rename"
	"github.com/spf13/cobra"
)

var FlagDry bool
var FlagVersion string

var Cmd = &cobra.Command{
	Use:     "spaces-and-commas",
	Short:   "Replace spaces and commans in filenames with underscores",
	Aliases: []string{"sac"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		replaceSpacesAndCommas()
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}

func replaceSpacesAndCommas() {
	dir := "."

	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {
		if !file.IsDir() {
			oldName := file.Name()
			newName := strings.ReplaceAll(oldName, " - ", "__")
			newName = strings.ReplaceAll(newName, ",", "_")
			newName = strings.ReplaceAll(newName, " ", "_")

			if oldName != newName {
				oldPath := filepath.Join(dir, oldName)
				newPath := filepath.Join(dir, newName)

				err = os.Rename(oldPath, newPath)
				if err != nil {
					fmt.Printf("Error renaming file '%s' to '%s': %s\n", oldName, newName, err)
				} else {
					fmt.Printf("Renamed '%s' to '%s'\n", oldName, newName)
				}
			}
		}
	}
}
