package elastic_password

import (
	"encoding/json"
	"fmt"

	parent_cmd "github.com/sikalabs/slu/cmd/eck"
	"github.com/sikalabs/slu/utils/eck_utils"

	"github.com/spf13/cobra"
)

var FlagJson bool
var FlagNamespace string
var FlagElasticsearchName string

var Cmd = &cobra.Command{
	Use:     "elastic-password",
	Short:   "Get Passowrd for Elastic User",
	Aliases: []string{"ep", "password", "p"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		password := eck_utils.GetElasticPassword(FlagNamespace, FlagElasticsearchName)

		if FlagJson {
			outJson, err := json.Marshal(password)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(outJson))
		} else {
			fmt.Println(password)
		}
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagNamespace,
		"namespace",
		"n",
		"",
		"ElasticSearch Namespace",
	)
	Cmd.MarkFlagRequired("namespace")
	Cmd.Flags().StringVarP(
		&FlagElasticsearchName,
		"elasticsearch-name",
		"e",
		"",
		"ElasticSearch Name",
	)
	Cmd.MarkFlagRequired("elasticsearch-name")
	Cmd.PersistentFlags().BoolVar(
		&FlagJson,
		"json",
		false,
		"Format output to JSON",
	)
}
