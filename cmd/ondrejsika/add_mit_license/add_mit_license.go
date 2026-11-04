package add_mit_license

import (
	"fmt"
	"os"
	"time"

	"github.com/sikalabs/slu/cmd/ondrejsika"
	"github.com/sikalabs/slu/utils/exec_utils"
	"github.com/spf13/cobra"
)

func init() {
	ondrejsika.Cmd.AddCommand(Cmd)
}

var Cmd = &cobra.Command{
	Use:   "add-mit-license",
	Short: "Add MIT license by Ondrej Sika & SikaLabs",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		addMitLicense()
	},
}

func addMitLicense() {
	year := time.Now().Year()
	license := fmt.Sprintf(`MIT License

Copyright (c) %d Ondrej Sika & SikaLabs

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
`, year)
	err := os.WriteFile("LICENSE", []byte(license), 0644)
	if err != nil {
		panic(err)
	}
	exec_utils.ExecOut("git", "add", "LICENSE")
	exec_utils.ExecOut("git", "commit", "-m", "feat(LICENSE): Add MIT license")
}
