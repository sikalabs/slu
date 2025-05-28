package do_sl_training_otp

import (
	"fmt"
	"log"
	"time"

	"github.com/pquerna/otp/totp"
	parent_cmd "github.com/sikalabs/slu/cmd/scripts"
	"github.com/spf13/cobra"
)

var FlagDry bool

var Cmd = &cobra.Command{
	Use:   "do-sl-training-otp",
	Short: "Get OTP for read-only user in SikaLabs Training DigitalOcean",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		otp, err := totp.GenerateCode("V7OUDK7CSY2Y5NACGPIUSRUONQNEHR45", time.Now())
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(otp)
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}
