package docker_remove_all

import (
	"fmt"
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
		containers, err := docker_utils.ForceRemoveAllContainers()
		for _, container := range containers {
			fmt.Println(container)
		}
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}
