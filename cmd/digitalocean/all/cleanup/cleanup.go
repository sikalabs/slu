package cleanup

import (
	"fmt"
	"time"

	parent_cmd "github.com/sikalabs/slu/cmd/digitalocean/all"
	"github.com/sikalabs/slu/utils/digitalocean_utils"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "cleanup",
	Short: "Clean up all not used resources (only volumes yet)",
	Long: `Clean up all not used resources

	Currently, all resources means:

	- Volumes
	`,
	Args: cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		token := digitalocean_utils.GetToken()

		// Prepare clean ups
		v := digitalocean_utils.PrepareVolumesCleanUp(token)

		fmt.Println("")
		fmt.Println("Wait 10s, cancel clean up using ctrl-c")

		for i := 1; i <= 10; i++ {
			fmt.Print(".")
			time.Sleep(time.Second)
		}

		fmt.Println("")
		fmt.Println("Cleaning up...")

		// Do clean ups
		digitalocean_utils.DoVolumesCleanUp(token, v)

		fmt.Println("Done.")
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}
