package add_charts

import (
	"path/filepath"
	"strings"

	commit_cmd "github.com/sikalabs/slu/cmd/git/commit"
	"github.com/sikalabs/slu/utils/exec_utils"
	"github.com/sikalabs/slu/utils/git_utils"
	"github.com/spf13/cobra"
)

func splitByLast(s string, sep string) (string, string, error) {
	lastInd := strings.LastIndex(s, sep)
	return s[:lastInd], s[lastInd+1:], nil
}

func getPackageFromFilename(sourceName string) (string, string, error) {
	tmp := filepath.Base(sourceName)
	tmp = strings.ReplaceAll(tmp, ".tgz", "")
	name, version, _ := splitByLast(tmp, "-")
	return name, version, nil
}

var Cmd = &cobra.Command{
	Use:     "add-charts",
	Short:   "Commit change: Add Charts (.tgz) to Helm Git repo",
	Aliases: []string{},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		allNewFiles, _ := git_utils.GetNewAddedFiles(".")
		var messageFilesArray []string
		for _, x := range allNewFiles {
			// All packages ends with .tgz
			if !strings.HasSuffix(x, ".tgz") {
				continue
			}
			name, version, _ := getPackageFromFilename(x)
			messageFilesArray = append(messageFilesArray, name+" "+version)
		}
		messageFilesString := strings.Join(messageFilesArray, ", ")
		exec_utils.ExecOut("git", "add", ".")
		exec_utils.ExecOut(
			"git", "commit", "-n", "-m",
			"[auto] feat: Add Charts ("+messageFilesString+")")
	},
}

func init() {
	commit_cmd.Cmd.AddCommand(Cmd)
}
