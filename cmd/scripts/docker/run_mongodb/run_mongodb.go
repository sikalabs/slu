package run_mongodb

import (
	"fmt"
	"strings"

	parent_cmd "github.com/sikalabs/slu/cmd/scripts/docker"
	"github.com/sikalabs/slu/utils/exec_utils"
	"github.com/spf13/cobra"
)

var FlagVolume string
var FlagUsername string
var FlagPassword string
var FlagDryRun bool

var Cmd = &cobra.Command{
	Use:     "run-mongodb",
	Short:   "docker run --name mongodb -d -p 27017:27017 mongo",
	Aliases: []string{"rmongo", "rmongodb"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		dockerArgs := []string{
			"run",
			"--name", "mongodb",
			"-d",
			"-p", "27017:27017",
		}

		if FlagVolume != "" {
			dockerArgs = append(dockerArgs, "-v", FlagVolume+":/data/db")
		}

		if FlagPassword != "" {
			username := FlagUsername
			if username == "" {
				username = "admin"
			}
			dockerArgs = append(dockerArgs,
				"-e", "MONGO_INITDB_ROOT_USERNAME="+username,
				"-e", "MONGO_INITDB_ROOT_PASSWORD="+FlagPassword,
			)
		}

		dockerArgs = append(dockerArgs, "mongo")

		if FlagDryRun {
			if FlagVolume != "" {
				fmt.Printf("docker volume create %s\n", FlagVolume)
			}
			fmt.Printf("docker %s\n", strings.Join(dockerArgs, " "))
			return
		}

		if FlagVolume != "" {
			exec_utils.ExecOut("docker", "volume", "create", FlagVolume)
		}

		exec_utils.ExecOut("docker", dockerArgs...)
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagVolume,
		"volume",
		"v",
		"",
		"Create and mount a Docker volume for MongoDB data",
	)
	Cmd.Flags().StringVarP(
		&FlagUsername,
		"username",
		"u",
		"",
		"Set MongoDB root username (default \"admin\", used only if --password is set)",
	)
	Cmd.Flags().StringVarP(
		&FlagPassword,
		"password",
		"p",
		"",
		"Set MongoDB root password (MONGO_INITDB_ROOT_PASSWORD)",
	)
	Cmd.Flags().BoolVar(
		&FlagDryRun,
		"dry-run",
		false,
		"Print command instead of running it",
	)
}
