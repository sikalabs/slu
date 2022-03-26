package required_environment_variables

import (
	"fmt"
	"os"

	parent_cmd "github.com/sikalabs/slu/cmd/shell_scripts"
	"github.com/spf13/cobra"
)

var FlagEnvVar []string

var Cmd = &cobra.Command{
	Use:     "required-environment-variables",
	Short:   "Validate required envirnonment variables",
	Aliases: []string{"req-env-var", "rev"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		ok := true
		for _, envName := range FlagEnvVar {
			env := os.Getenv(envName)
			if env == "" {
				ok = false
				fmt.Printf("Environment variable \"%s\" is missing or blank\n", envName)
			}
		}
		if !ok {
			os.Exit(1)
		}
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringArrayVarP(
		&FlagEnvVar,
		"env-var",
		"e",
		[]string{},
		"Environment variable",
	)
	Cmd.MarkFlagRequired("env-var")
}
