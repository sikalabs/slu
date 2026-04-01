package gb2mb

import (
	"fmt"
	"strconv"

	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "gb2mb <gb>",
	Short: "Convert gigabytes to megabytes",
	Args:  cobra.ExactArgs(1),
	Run: func(c *cobra.Command, args []string) {
		gb, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Error: invalid number")
			return
		}
		fmt.Println(gb * 1024)
	},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
