package list

import (
	"fmt"
	"io/ioutil"
	"log"

	registry_client "github.com/heroku/docker-registry-client/registry"
	parent_cmd "github.com/sikalabs/slu/cmd/docker/registry"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "list <registry url>",
	Short:   "List Images",
	Args:    cobra.ExactArgs(1),
	Aliases: []string{"ls"},
	Run: func(c *cobra.Command, args []string) {
		registry := args[0]
		url := "https://" + registry
		username := ""
		password := ""
		registry_client.Quiet("")
		log.SetOutput(ioutil.Discard)
		r, _ := registry_client.New(url, username, password)
		r.Logf = registry_client.Quiet
		repositories, _ := r.Repositories()
		for _, image := range repositories {
			fmt.Println(registry + "/" + image)
		}
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}
