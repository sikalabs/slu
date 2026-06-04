package mit_license

import (
	"fmt"
	"os"
	"time"

	file_templates_cmd "github.com/sikalabs/slu/cmd/file_templates"
	"github.com/spf13/cobra"
)

var FlagOndrejSika bool
var FlagSikaLabs bool
var FlagAll bool

var Cmd = &cobra.Command{
	Use:   "mit-license",
	Short: "Create MIT LICENSE file",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		if FlagAll {
			FlagOndrejSika = true
			FlagSikaLabs = true
		}

		if !FlagOndrejSika && !FlagSikaLabs {
			fmt.Fprintln(os.Stderr, "Error: at least one of --ondrejsika or --sikalabs is required")
			os.Exit(1)
		}

		var authors []string
		if FlagOndrejSika {
			authors = append(authors, "Ondrej Sika")
		}
		if FlagSikaLabs {
			authors = append(authors, "SikaLabs s.r.o.")
		}

		var authorStr string
		if len(authors) == 2 {
			authorStr = authors[0] + " and " + authors[1]
		} else {
			authorStr = authors[0]
		}

		year := time.Now().Year()
		content := fmt.Sprintf(`MIT License

Copyright (c) %d %s

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the “Software”), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
`, year, authorStr)

		err := os.WriteFile("LICENSE", []byte(content), 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing LICENSE: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	file_templates_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().BoolVar(&FlagOndrejSika, "ondrejsika", false, "Use Ondrej Sika as author")
	Cmd.Flags().BoolVar(&FlagSikaLabs, "sikalabs", false, "Use SikaLabs s.r.o. as author")
	Cmd.Flags().BoolVarP(&FlagAll, "all", "a", false, "Use both Ondrej Sika and SikaLabs s.r.o. as authors")
}
