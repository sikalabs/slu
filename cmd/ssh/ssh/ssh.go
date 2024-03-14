package ssh

import (
	"fmt"
	"log"
	"strings"

	"github.com/k0sproject/rig"
	parent_cmd "github.com/sikalabs/slu/cmd/ssh"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const DEFAULT_KEY_LENGTH = 2048

var FlagUseGo bool
var FlagECDSA bool

var Cmd = &cobra.Command{
	Use:   "ssh <user@host> <command> [args...]",
	Short: "Connect to SSH server",
	Args:  cobra.MinimumNArgs(2),
	Run: func(c *cobra.Command, args []string) {
		user, host, err := parseSSHUserAndHost(args[0])
		if err != nil {
			log.Fatal(err)
		}
		ssh(user, host, args[1:])
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}

func ssh(user, address string, args []string) {
	logger := logrus.New()
	rig.SetLogger(logger)

	conn := rig.Connection{
		SSH: &rig.SSH{
			User:    user,
			Address: address,
			// PasswordCallback: func() (string, error) {
			// 	return "password", nil
			// },
		},
	}
	if err := conn.Connect(); err != nil {
		logger.Fatal(err)
	}
	defer conn.Disconnect()

	output, err := conn.ExecOutput(strings.Join(args, " "))
	if err != nil {
		logger.Fatal(err)
	}
	fmt.Println(output)
}

func parseSSHUserAndHost(sshString string) (user, host string, err error) {
	parts := strings.Split(sshString, "@")
	if len(parts) == 1 {
		user = "root"
		host = parts[0]
		return
	}
	if len(parts) == 2 {
		user = parts[0]
		host = parts[1]
		return
	}

	err = fmt.Errorf("invalid SSH string format, must be host or user@host")
	return
}
