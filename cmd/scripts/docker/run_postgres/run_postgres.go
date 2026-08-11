package run_postgres

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
var FlagDatabase string
var FlagDryRun bool

var Cmd = &cobra.Command{
	Use:     "run-postgres",
	Short:   "docker run --name postgres -d -p 5432:5432 -e ... postgres",
	Aliases: []string{"rpg", "rpostgres"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		dockerArgs := []string{
			"run",
			"--name", "postgres",
			"-d",
			"-p", "5432:5432",
		}

		if FlagVolume != "" {
			dockerArgs = append(dockerArgs, "-v", FlagVolume+":/var/lib/postgresql/data")
		}

		username := FlagUsername
		if username == "" {
			username = "postgres"
		}
		dockerArgs = append(dockerArgs, "-e", "POSTGRES_USER="+username)

		password := FlagPassword
		if password == "" {
			password = "pg"
		}
		dockerArgs = append(dockerArgs, "-e", "POSTGRES_PASSWORD="+password)

		if FlagDatabase != "" {
			dockerArgs = append(dockerArgs, "-e", "POSTGRES_DB="+FlagDatabase)
		}

		dockerArgs = append(dockerArgs, "postgres")

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
		"Create and mount a Docker volume for PostgreSQL data",
	)
	Cmd.Flags().StringVarP(
		&FlagUsername,
		"username",
		"u",
		"",
		"Set PostgreSQL username (default \"postgres\", POSTGRES_USER)",
	)
	Cmd.Flags().StringVarP(
		&FlagPassword,
		"password",
		"p",
		"",
		"Set PostgreSQL password (default \"pg\", POSTGRES_PASSWORD)",
	)
	Cmd.Flags().StringVarP(
		&FlagDatabase,
		"database",
		"d",
		"",
		"Create a database with this name (POSTGRES_DB)",
	)
	Cmd.Flags().BoolVar(
		&FlagDryRun,
		"dry-run",
		false,
		"Print command instead of running it",
	)
}
