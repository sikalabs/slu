package root

import (
	"github.com/sikalabs/slu/version"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "slu",
	Short: "SikaLabs Utils, " + version.Version,
}
