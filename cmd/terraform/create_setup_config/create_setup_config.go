package create_setup_config

import (
	"encoding/json"
	"os"
	"path/filepath"

	parentcmd "github.com/sikalabs/slu/cmd/terraform"
	"github.com/sikalabs/slu/internal/error_utils"
	"github.com/spf13/cobra"
)

type TerraformConfig struct {
	Meta struct {
		SchemaVersion string `json:"SchemaVersion"`
	} `json:"Meta"`
	GitlabURL  string `json:"GitlabURL"`
	ProjectID  string `json:"ProjectID"`
	StateName  string `json:"StateName"`
}

var Cmd = &cobra.Command{
	Use:   "create-setup-config <gitlab-url> <project-id> <state-name>",
	Short: "Create Terraform setup config file",
	Args:  cobra.ExactArgs(3),
	Run: func(c *cobra.Command, args []string) {
		gitlabURL := args[0]
		projectID := args[1]
		stateName := args[2]

		// Create config structure
		config := TerraformConfig{
			GitlabURL: gitlabURL,
			ProjectID: projectID,
			StateName: stateName,
		}
		config.Meta.SchemaVersion = "1"

		// Create directory
		configDir := filepath.Join(".sikalabs", "terraform")
		err := os.MkdirAll(configDir, 0755)
		error_utils.HandleError(err, "Failed to create directory")

		// Marshal to JSON with indentation
		jsonData, err := json.MarshalIndent(config, "", "  ")
		error_utils.HandleError(err, "Failed to marshal JSON")

		// Write to file
		configPath := filepath.Join(configDir, "terraform.json")
		err = os.WriteFile(configPath, jsonData, 0644)
		error_utils.HandleError(err, "Failed to write config file")
	},
}

func init() {
	parentcmd.Cmd.AddCommand(Cmd)
}
