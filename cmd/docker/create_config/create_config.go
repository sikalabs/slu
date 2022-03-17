package create_config

import (
	"encoding/base64"
	"fmt"

	parent_cmd "github.com/sikalabs/slu/cmd/docker"
	"github.com/spf13/cobra"
)

var FlagRegistry string
var FlagRegistryUsername string
var FlagRegistryPassword string

var Cmd = &cobra.Command{
	Use:     "create-config",
	Short:   "Create .docker/config.json (for Kubernetes)",
	Aliases: []string{"create-conf", "cc"},
	Run: func(c *cobra.Command, args []string) {
		auth := FlagRegistryUsername + ":" + FlagRegistryPassword
		authBase64 := base64.StdEncoding.EncodeToString([]byte(auth))
		fmt.Printf(`{"auths":{"%s":{"auth":"%s"}}`, FlagRegistry, authBase64)
		fmt.Printf("\n")
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagRegistry,
		"registry",
		"r",
		"",
		"Registry (eg.: registry.sikalabs.com)",
	)
	Cmd.MarkFlagRequired("registry")
	Cmd.Flags().StringVarP(
		&FlagRegistryUsername,
		"username",
		"u",
		"",
		"Username",
	)
	Cmd.MarkFlagRequired("username")
	Cmd.Flags().StringVarP(
		&FlagRegistryPassword,
		"password",
		"p",
		"",
		"Password",
	)
	Cmd.MarkFlagRequired("password")
}
