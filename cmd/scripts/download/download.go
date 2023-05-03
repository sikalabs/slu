package download

import (
	"io"
	"log"
	"net/http"
	"os"

	parent_cmd "github.com/sikalabs/slu/cmd/scripts"
	"github.com/sikalabs/slu/utils/exec_utils"
	"github.com/spf13/cobra"
)

var FlagUseWget bool
var FlagUseCurl bool

var Cmd = &cobra.Command{
	Use:   "download",
	Short: "Download Something",
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
	Cmd.PersistentFlags().BoolVarP(
		&FlagUseWget,
		"wget",
		"w",
		false,
		"Download using wget",
	)
	Cmd.PersistentFlags().BoolVarP(
		&FlagUseCurl,
		"curl",
		"c",
		false,
		"Download using curl",
	)
	Cmd.MarkFlagsMutuallyExclusive("wget", "curl")
	buildCmd(
		Cmd,
		"https://cdimage.debian.org/debian-cd/current/amd64/iso-cd/debian-11.7.0-amd64-netinst.iso",
		"debian-11.7.0-amd64-netinst.iso",
		"debian-iso",
		[]string{},
	)
}

func buildCmd(cmd *cobra.Command, url string, outputFileName string, name string, aliases []string) {
	var Cmd = &cobra.Command{
		Use:     name,
		Short:   "Download " + name,
		Aliases: aliases,
		Args:    cobra.NoArgs,
		Run: func(c *cobra.Command, args []string) {
			if FlagUseCurl {
				exec_utils.ExecOut("curl", "-L", "-o", outputFileName, url)
			} else if FlagUseWget {
				exec_utils.ExecOut("wget", "-O", outputFileName, url)
			} else {
				webToBin(url, outputFileName)
			}
		},
	}
	cmd.AddCommand(Cmd)
}

func webToBin(url, outFileName string) {
	var err error

	resp, err := http.Get(url)
	handleError(err)
	defer resp.Body.Close()

	outFile, err := os.OpenFile(outFileName, os.O_CREATE|os.O_WRONLY, 0755)
	handleError(err)
	defer outFile.Close()

	_, err = io.Copy(outFile, resp.Body)
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
