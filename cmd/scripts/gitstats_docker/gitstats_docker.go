package gitstats_docker

import (
	"fmt"
	"log"
	"os"
	"time"

	parent_cmd "github.com/sikalabs/slu/cmd/scripts"
	"github.com/sikalabs/slu/utils/exec_utils"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "gitstats-docker",
	Short:   "Generate GitStats (using Docker)",
	Aliases: []string{"gitstats", "gs"},
	Run: func(c *cobra.Command, args []string) {
		workdir, err := os.Getwd()
		if err != nil {
			log.Fatalln(err)
		}
		containerName := fmt.Sprintf("gitstats-%d", time.Now().Unix())
		fmt.Println("Generating GitStats in Docker container: " + containerName + " ...")
		exec_utils.ExecNoOut(
			"docker", "run", "-d",
			"--name", containerName,
			"-w", "/workspace",
			"-v", workdir+"/.git:/workspace/.git:ro",
			"-v", workdir+"/.gitstats_output:/workspace/.gitstats_output:rw",
			"nixery.dev/git/gitstats",
			"gitstats", ".", ".gitstats_output",
		)
		exec_utils.ExecNoOut(
			"docker", "wait", containerName,
		)
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}
