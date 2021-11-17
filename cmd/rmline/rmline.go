package version

import (
	"log"
	"strconv"
	"strings"

	"github.com/sikalabs/slu/cmd/root"
	"github.com/sikalabs/slu/utils/file_utils"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "rmline </home/ondrej/file:1>",
	Short: "Remove line in file",
	Args:  cobra.ExactArgs(1),
	Run: func(c *cobra.Command, args []string) {
		s := strings.Split(args[0], ":")
		if len(s) != 2 {
			log.Fatal("Wrong input format")
		}
		filename := s[0]
		line, err := strconv.Atoi(s[1])
		if err != nil {
			log.Fatal(err)
		}
		err = file_utils.RemoveLines(filename, line, 1)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
