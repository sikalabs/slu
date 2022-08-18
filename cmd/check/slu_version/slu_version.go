package password_reset

import (
	"log"

	parent_cmd "github.com/sikalabs/slu/cmd/check"
	"github.com/sikalabs/slu/utils/semver_utils"
	"github.com/sikalabs/slu/version"

	"github.com/spf13/cobra"
)

var FlagVersion string

var Cmd = &cobra.Command{
	Use:   "slu-version",
	Short: "slu version must be higher or equal",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		ok := semver_utils.CheckMinimumVersion(version.Version, FlagVersion)
		if !ok {
			log.Fatal("slu version (" + version.Version + ") must be higher or equal to " + FlagVersion)
		}
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagVersion,
		"version",
		"v",
		"",
		"slu version",
	)
	Cmd.MarkFlagRequired("version")
}
