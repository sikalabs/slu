package cmd

import (
	_ "github.com/sikalabs/slut/cmd/mysql"
	_ "github.com/sikalabs/slut/cmd/mysql/create"
	_ "github.com/sikalabs/slut/cmd/mysql/drop"
	"github.com/sikalabs/slut/cmd/root"
	_ "github.com/sikalabs/slut/cmd/version"
	"github.com/spf13/cobra"
)

func Execute() {
	cobra.CheckErr(root.RootCmd.Execute())
}
