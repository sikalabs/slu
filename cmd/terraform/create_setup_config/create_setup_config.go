package create_setup_config

import (
	"encoding/json"
	"os"
	"path/filepath"

	parentcmd "github.com/sikalabs/slu/cmd/terraform"
	"github.com/sikalabs/slu/internal/error_utils"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type TerraformConfig struct {
	Meta struct {
		SchemaVersion string `json:"SchemaVersion" yaml:"SchemaVersion"`
	} `json:"Meta" yaml:"Meta"`
	GitlabURL  string `json:"GitlabURL" yaml:"GitlabURL"`
	ProjectID  string `json:"ProjectID" yaml:"ProjectID"`
	StateName  string `json:"StateName" yaml:"StateName"`
}

var FlagFormat string

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

		var data []byte
		var fileName string

		// Marshal based on format
		if FlagFormat == "yaml" {
			data, err = yaml.Marshal(config)
			error_utils.HandleError(err, "Failed to marshal YAML")
			fileName = "terraform.yaml"
		} else {
			data, err = json.MarshalIndent(config, "", "  ")
			error_utils.HandleError(err, "Failed to marshal JSON")
			fileName = "terraform.json"
		}

		// Write to file
		configPath := filepath.Join(configDir, fileName)
		err = os.WriteFile(configPath, data, 0644)
		error_utils.HandleError(err, "Failed to write config file")
	},
}

func init() {
	parentcmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagFormat,
		"format",
		"f",
		"yaml",
		"Output format: json or yaml (default: yaml)",
	)
}
