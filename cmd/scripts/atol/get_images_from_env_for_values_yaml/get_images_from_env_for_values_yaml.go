package get_images_from_env_for_values_yaml

import (
	"fmt"
	"strings"

	"github.com/joho/godotenv"
	parent_cmd "github.com/sikalabs/slu/cmd/scripts/atol"
	"github.com/sikalabs/slu/internal/error_utils"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "get-images-from-env-for-values-yaml",
	Short: "Get images from .env for values.yaml",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		GetImagesFromEnvForValuesYaml()
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}

func GetImagesFromEnvForValuesYaml() {
	dotEnvVars, err := godotenv.Read(".env")
	error_utils.HandleError(err, "Failed to read .env file")

	out := []string{}

	for name, value := range dotEnvVars {
		if strings.Contains(name, "IMAGE") {
			out = append(out, fmt.Sprintf("%s=%s", name, value))
		}
	}

	for _, line := range out {
		fmt.Println(line)
	}
}
