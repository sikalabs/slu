package version

import (
	"strings"

	parent_cmd "github.com/sikalabs/slu/cmd/check"
	"github.com/sikalabs/slu/internal/error_utils"
	"github.com/sikalabs/slu/utils/exec_utils"
	"github.com/sikalabs/slu/utils/semver_utils"

	"github.com/spf13/cobra"
)

var FlagBinary string
var FlagVersion string

var Cmd = &cobra.Command{
	Use:   "version",
	Short: "version must be higher or equal",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		currentVersionRaw, err := exec_utils.ExecStr(FlagBinary, "version")
		error_utils.HandleError(err, "Failed to get version")
		currentVersion := strings.ReplaceAll(currentVersionRaw, "\n", "")
		ok := semver_utils.CheckMinimumVersion(currentVersion, FlagVersion)
		error_utils.HandleNotOK(ok, FlagBinary+" version ("+currentVersion+") must be higher or equal to "+FlagVersion)
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagBinary,
		"binary",
		"b",
		"",
		"binary",
	)
	Cmd.MarkFlagRequired("binary")
	Cmd.Flags().StringVarP(
		&FlagVersion,
		"version",
		"v",
		"",
		"version",
	)
	Cmd.MarkFlagRequired("version")
}
