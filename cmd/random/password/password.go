package password

import (
	"fmt"
	"log"
	"strings"
	"unicode"

	random_cmd "github.com/sikalabs/slu/cmd/random"
	"github.com/sikalabs/slu/utils/random_utils"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "password",
	Short:   "Generate random password",
	Aliases: []string{"pwd", "passwd", "pass"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		i := 0
		for {
			s := random_utils.RandomString(16, random_utils.ALL)
			if containsLowercase(s) && containsUpercase(s) && containsDigit(s) {
				fmt.Println(addUnderscores(s))
				break
			}
			if i > 20 {
				log.Fatalln("Cannot generate password")
			}
		}
	},
}

func init() {
	random_cmd.Cmd.AddCommand(Cmd)
}

func containsLowercase(s string) bool {
	for _, r := range s {
		if unicode.IsLower(r) {
			return true
		}
	}
	return false
}

func containsUpercase(s string) bool {
	for _, r := range s {
		if unicode.IsLower(r) {
			return true
		}
	}
	return false
}

func containsDigit(s string) bool {
	for _, r := range s {
		if unicode.IsDigit(r) {
			return true
		}
	}
	return false
}

func addUnderscores(input string) string {
	var blocks []string
	for i := 0; i < len(input); i += 4 {
		end := i + 4
		if end > len(input) {
			end = len(input)
		}
		blocks = append(blocks, input[i:end])
	}

	return strings.Join(blocks, "_")
}
