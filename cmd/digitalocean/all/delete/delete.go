package delete

import (
	"fmt"
	"time"

	parent_cmd "github.com/sikalabs/slu/cmd/digitalocean/all"
	"github.com/sikalabs/slu/utils/digitalocean_utils"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete ALL resources in DigitalOcean account",
	Long: `Delete ALL resources in DigitalOcean account

	Currently, all resources means:

	- Droplets
	- SSH Keys
	`,
	Args: cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		token := digitalocean_utils.GetToken()

		// Prepare Delete
		droplets := digitalocean_utils.PrepareAllDropletsDelete(token)
		sshKeys := digitalocean_utils.PrepareAllSSHKeysDelete(token)

		fmt.Println("")
		fmt.Println("Wait 10s, cancel clean up using ctrl-c")

		for i := 1; i <= 10; i++ {
			fmt.Print(".")
			time.Sleep(time.Second)
		}

		fmt.Println("")
		fmt.Println("")
		fmt.Println("Delete ...")
		fmt.Println("")

		// Do Delete
		digitalocean_utils.DoAllDropletsDelete(token, droplets)
		digitalocean_utils.DoAllSSHKeysDelete(token, sshKeys)

		fmt.Println("Done.")
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}
