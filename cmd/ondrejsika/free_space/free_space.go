package free_space

import (
	"fmt"
	"log"

	parentcmd "github.com/sikalabs/slu/cmd/ondrejsika"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "free-space",
	Short: "Print free space on root disk (/)",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		printFreeSpace()
	},
}

func init() {
	parentcmd.Cmd.AddCommand(Cmd)
}

func toGB(b uint64) float64 {
	return float64(b) / float64(1<<30)
}

func printFreeSpace() {
	freeSpace := getFreeSpaceOrDie()
	fmt.Printf("%.2f GB\n", toGB(freeSpace))
}

func handleError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
