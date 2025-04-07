package latest_version_go

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	parent_cmd "github.com/sikalabs/slu/cmd/scripts"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "latest-version-go <script>",
	Aliases: []string{"lvgo"},
	Short:   "Get latest version of Go from https://go.dev/dl/?mode=json",
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		fmt.Println(getLatestVersionOfGoOrDie())
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}

type GoVersion struct {
	Version string `json:"version"`
}

func getLatestVersionOfGoOrDie() string {
	resp, err := http.Get("https://go.dev/dl/?mode=json")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	var versions []GoVersion
	if err := json.NewDecoder(resp.Body).Decode(&versions); err != nil {
		log.Fatalln(err)
	}

	if len(versions) == 0 {
		log.Fatalln(err)
	}

	latest := versions[0].Version
	return strings.TrimPrefix(latest, "go")
}
