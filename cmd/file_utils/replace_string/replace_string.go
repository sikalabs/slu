package replace_string

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	parent_cmd "github.com/sikalabs/slu/cmd/file_utils"

	"github.com/spf13/cobra"
)

var FlagInputFilePath string
var FlagOutputFilePath string
var FlagFindString string
var FlagReplacementString string

var Cmd = &cobra.Command{
	Use:     "replace-string",
	Short:   "Replace string in file (like sed)",
	Aliases: []string{"rs"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		inputBytes, err := ioutil.ReadFile(FlagInputFilePath)
		if err != nil {
			log.Fatalln(err)
		}

		input := string(inputBytes)
		output := strings.Replace(input, FlagFindString, FlagReplacementString, -1)

		if FlagOutputFilePath == "" {
			fmt.Println(output)
		} else {
			err = ioutil.WriteFile(FlagOutputFilePath, []byte(output), 0644)
			if err != nil {
				log.Fatalln(err)
			}
		}
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagInputFilePath,
		"input-file",
		"i",
		"",
		"Input file path",
	)
	Cmd.MarkFlagRequired("input-file")
	Cmd.Flags().StringVarP(
		&FlagOutputFilePath,
		"output-file",
		"o",
		"",
		"Output file path (default is STDOUT)",
	)
	Cmd.Flags().StringVarP(
		&FlagFindString,
		"find",
		"f",
		"",
		"Find string",
	)
	Cmd.MarkFlagRequired("find")
	Cmd.Flags().StringVarP(
		&FlagReplacementString,
		"replacement",
		"r",
		"",
		"Replacement string",
	)
	Cmd.MarkFlagRequired("replacement")
}
