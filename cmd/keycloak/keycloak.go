package keycloak

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "keycloak",
	Short: "Keycloak Utils",
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
