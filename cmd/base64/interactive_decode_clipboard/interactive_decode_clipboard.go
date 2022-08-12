package interactive_decode_clipboard

import (
	"encoding/base64"
	"fmt"
	"log"
	"time"

	"github.com/atotto/clipboard"
	parent_cmd "github.com/sikalabs/slu/cmd/base64"

	"github.com/spf13/cobra"
)

var FlagNamespace string

var Cmd = &cobra.Command{
	Use:     "interactive-decode-clipboard",
	Short:   "Base64 Interactive Decode from Clipboard",
	Aliases: []string{"idc"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		clipboardDataOld := ""
		for {
			clipboardData, err := clipboard.ReadAll()
			if err != nil {
				log.Fatal(err)
			}
			if clipboardData != clipboardDataOld {
				clipboardDataOld = clipboardData
				dataPlainText, err := base64.StdEncoding.DecodeString(clipboardData)
				if err == nil {
					fmt.Println(string(dataPlainText))
					fmt.Println("----------------------------------------------------")
				}
			}
			time.Sleep(1 * time.Second)
		}
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
}
