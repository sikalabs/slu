package json_to_yaml

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	parentcmd "github.com/sikalabs/slu/cmd/scripts"
	"github.com/sikalabs/slu/internal/error_utils"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var Cmd = &cobra.Command{
	Use:   "json-to-yaml <json-file>",
	Short: "Convert JSON file to YAML in place",
	Long: `Convert a JSON file to YAML format.
Creates a new .yaml file and removes the original .json file.

Example: slu scripts json-to-yaml example.json
  Creates: example.yaml
  Removes: example.json`,
	Args: cobra.ExactArgs(1),
	Run: func(c *cobra.Command, args []string) {
		jsonFilePath := args[0]

		// Read JSON file
		jsonData, err := os.ReadFile(jsonFilePath)
		error_utils.HandleError(err, "Failed to read JSON file:")

		// Unmarshal JSON to generic interface
		var data interface{}
		err = json.Unmarshal(jsonData, &data)
		error_utils.HandleError(err, "Failed to parse JSON file:")

		// Marshal to YAML
		yamlData, err := yaml.Marshal(data)
		error_utils.HandleError(err, "Failed to convert to YAML:")

		// Determine YAML file path
		ext := filepath.Ext(jsonFilePath)
		var yamlFilePath string
		if ext == ".json" {
			yamlFilePath = strings.TrimSuffix(jsonFilePath, ext) + ".yaml"
		} else {
			// If file doesn't have .json extension, just append .yaml
			yamlFilePath = jsonFilePath + ".yaml"
		}

		// Write YAML file
		err = os.WriteFile(yamlFilePath, yamlData, 0644)
		error_utils.HandleError(err, "Failed to write YAML file:")

		// Remove original JSON file
		err = os.Remove(jsonFilePath)
		error_utils.HandleError(err, "Failed to remove JSON file:")
	},
}

func init() {
	parentcmd.Cmd.AddCommand(Cmd)
}
