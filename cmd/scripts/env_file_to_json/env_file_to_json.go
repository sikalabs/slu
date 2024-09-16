package env_file_to_json

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	parent_cmd "github.com/sikalabs/slu/cmd/scripts"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "env-file-to-json",
	Short: "Transform .env file to JSON for HashiCorp Vault",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		envFileToJson()
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}

func envFileToJson() {
	// Load .env file into a map
	envMap, err := godotenv.Read(".env")
	if err != nil {
		log.Fatalf("Error reading .env file: %v", err)
	}

	// Convert the map to pretty-printed JSON
	jsonData, err := json.MarshalIndent(envMap, "", "  ")
	if err != nil {
		log.Fatalf("Error converting to JSON: %v", err)
	}

	// Print the pretty-printed JSON
	fmt.Println(string(jsonData))
}
