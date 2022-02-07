package list

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	registry_client "github.com/heroku/docker-registry-client/registry"
	parent_cmd "github.com/sikalabs/slu/cmd/docker/registry"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "tags <image>",
	Short:   "List Tags of Image",
	Args:    cobra.ExactArgs(1),
	Aliases: []string{"ls"},
	Run: func(c *cobra.Command, args []string) {
		image := args[0]
		imageSplit := strings.SplitN(image, "/", 2)
		registry := imageSplit[0]
		path := imageSplit[1]
		url := "https://" + registry
		username := ""
		password := ""
		registry_client.Quiet("")
		log.SetOutput(ioutil.Discard)
		r, _ := registry_client.New(url, username, password)
		r.Logf = registry_client.Quiet
		repositories, _ := r.Tags(path)
		for _, tag := range repositories {
			fmt.Println(image + ":" + tag)
		}
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}
