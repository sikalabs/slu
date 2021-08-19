package root

import (
	"github.com/sikalabs/slu/version"
	"github.com/spf13/cobra"
)

var RootCmdFlagJson bool

var RootCmd = &cobra.Command{
	Use:   "slu",
	Short: "SikaLabs Utils, " + version.Version,
}

func init() {
	RootCmd.PersistentFlags().BoolVar(
		&RootCmdFlagJson,
		"json",
		false,
		"Formatu output to JSON",
	)
}
