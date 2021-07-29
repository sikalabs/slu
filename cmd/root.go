package cmd

import (
	"github.com/sikalabs/slut/cmd/version"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "slut",
	Short: "SikaLabs Utils",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.AddCommand(version.VersionCmd)
}
