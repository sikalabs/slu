package create_service

import (
	"fmt"
	"log"
	"os"

	parent_cmd "github.com/sikalabs/slu/cmd/systemd"
	"github.com/sikalabs/slu/utils/systemd_utils"

	"github.com/spf13/cobra"
)

var FlagCreateFile bool
var FlagName string
var FlagDescription string
var FlagUser string
var FlagGroup string
var FlagWorkingDirectory string
var FlagExecStart string

var Cmd = &cobra.Command{
	Use:     "create-service",
	Short:   "Create systemd service file",
	Aliases: []string{"cs"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		out, err := systemd_utils.CreateSystemdServiceString(
			FlagName,
			FlagDescription,
			FlagUser,
			FlagGroup,
			FlagWorkingDirectory,
			FlagExecStart,
		)
		if err != nil {
			log.Fatalln(err)
		}
		if FlagCreateFile {
			err := os.WriteFile(
				"/etc/systemd/system/"+FlagName+".service", []byte(out), 0644)
			if err != nil {
				log.Fatalln(err)
			}
		} else {
			fmt.Println(out)
		}
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().BoolVarP(
		&FlagCreateFile,
		"create-file",
		"c",
		false,
		"Create service file in /etc/systemd/system",
	)
	Cmd.Flags().StringVarP(
		&FlagName,
		"name",
		"n",
		"",
		"Systed service name",
	)
	Cmd.MarkFlagRequired("name")
	Cmd.Flags().StringVarP(
		&FlagDescription,
		"description",
		"d",
		"",
		"Systed service description",
	)
	Cmd.MarkFlagRequired("description")
	Cmd.Flags().StringVarP(
		&FlagUser,
		"user",
		"u",
		"",
		"Systed service user",
	)
	Cmd.MarkFlagRequired("user")
	Cmd.Flags().StringVarP(
		&FlagGroup,
		"group",
		"g",
		"",
		"Systed service group",
	)
	Cmd.MarkFlagRequired("group")
	Cmd.Flags().StringVarP(
		&FlagWorkingDirectory,
		"working-directory",
		"w",
		"",
		"Systed service working-directory",
	)
	Cmd.MarkFlagRequired("working-directory")
	Cmd.Flags().StringVarP(
		&FlagExecStart,
		"exec-start",
		"e",
		"",
		"Systed service exec-start",
	)
	Cmd.MarkFlagRequired("exec-start")
}
