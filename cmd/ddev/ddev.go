package ip

import (
	"log"
	"os"
	"os/exec"
	"strconv"

	"github.com/sikalabs/slu/cmd/root"
	"github.com/sikalabs/slu/utils/time_utils"
	"github.com/spf13/cobra"
)

var FlagVolume bool
var FlagImage string
var FlagShell string
var FlagHostNetwork bool

var Cmd = &cobra.Command{
	Use:   "ddev",
	Short: "Run sikalabs/dev in Docker",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		currentWorkDir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		cmdArgs := []string{
			"run",
			"--name", "dev-" + strconv.Itoa(time_utils.Unix()),
			"--rm", "-ti",
		}
		if FlagVolume {
			cmdArgs = append(
				cmdArgs,
				"-w", currentWorkDir,
				"-v", currentWorkDir+":"+currentWorkDir,
			)
		}
		cmdArgs = append(
			cmdArgs,
			FlagImage,
			FlagShell,
		)
		cmd := exec.Command("docker", cmdArgs...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		cmd.Run()
	},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
	Cmd.PersistentFlags().BoolVarP(
		&FlagVolume,
		"volume",
		"v",
		false,
		"Mount current directory to container",
	)
	Cmd.Flags().StringVarP(
		&FlagImage,
		"image",
		"i",
		"sikalabs/dev",
		"Container Image",
	)
	Cmd.Flags().BoolVar(
		&FlagHostNetwork,
		"host",
		false,
		"Use host network (--network=host)",
	)
	Cmd.Flags().StringVarP(
		&FlagShell,
		"shell",
		"s",
		"zsh",
		"Shell to run in container",
	)
}
