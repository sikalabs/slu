package otp

import (
	"fmt"
	"log"
	"time"

	"github.com/pquerna/otp/totp"
	parent_cmd "github.com/sikalabs/slu/cmd/scripts"
	"github.com/spf13/cobra"
)

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}

var Cmd = &cobra.Command{
	Use:   "otp <secret>",
	Short: "Generate OTP code from TOTP secret",
	Args:  cobra.ExactArgs(1),
	Run: func(c *cobra.Command, args []string) {
		code, err := totp.GenerateCode(args[0], time.Now())
		if err != nil {
			log.Fatal("Error generating OTP: ", err)
		}
		fmt.Printf("%s-%s\n", code[:3], code[3:])
	},
}
