package password

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"

	parent_cmd "github.com/sikalabs/slu/cmd/aws"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "who-am-i",
	Short:   "Get AWS organization",
	Aliases: []string{"w"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		printMasterAccountEmail()
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}

type Response struct {
	Organization struct {
		MasterAccountEmail string `json:"MasterAccountEmail"`
	} `json:"Organization"`
}

func printMasterAccountEmail() {
	out, err := exec.Command("aws", "organizations", "describe-organization").Output()
	if err != nil {
		log.Fatalln(err)
	}

	var response Response
	err = json.Unmarshal(out, &response)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(response.Organization.MasterAccountEmail)
}
