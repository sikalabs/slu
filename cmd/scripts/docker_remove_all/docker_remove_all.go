package docker_remove_all

import (
	"log"

	parent_cmd "github.com/sikalabs/slu/cmd/scripts"
	"github.com/sikalabs/slu/utils/docker_utils"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "docker-remove-all",
	Short:   "Docker Remove All Containers",
	Aliases: []string{"dra"},
	Run: func(c *cobra.Command, args []string) {
		err := docker_utils.ForceRemoveAllContainers()
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}
