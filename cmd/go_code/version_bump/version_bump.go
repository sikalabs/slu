package version_bump

import (
	"io/ioutil"

	"github.com/go-git/go-git/v5"
	go_code_cmd "github.com/sikalabs/slu/cmd/go_code"

	"github.com/spf13/cobra"
)

var CmdFlagVersion string
var CmdFlagNoCommit bool
var CmdFlagTag bool

var Cmd = &cobra.Command{
	Use:     "version-bump",
	Short:   "Bumb & commit version of (SL) Go application",
	Aliases: []string{"vb"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		version_go_file := `package version

var Version string = "` + CmdFlagVersion + `"
`
		err := ioutil.WriteFile("version/version.go", []byte(version_go_file), 0644)
		if err != nil {
			panic(err)
		}

		if CmdFlagNoCommit {
			return
		}

		r, err := git.PlainOpen(".")
		if err != nil {
			panic(err)
		}
		w, err := r.Worktree()
		if err != nil {
			panic(err)
		}
		_, err = w.Add("version/version.go")
		if err != nil {
			panic(err)
		}
		commit, _ := w.Commit("VERSION: "+CmdFlagVersion, &git.CommitOptions{})
		_, err = r.CommitObject(commit)
		if err != nil {
			panic(err)
		}

		if CmdFlagTag {
			_, err := r.CreateTag(
				CmdFlagVersion,
				commit,
				&git.CreateTagOptions{
					Message: "VERSION: " + CmdFlagVersion,
				},
			)
			if err != nil {
				panic(err)
			}
		}
	},
}

func init() {
	go_code_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&CmdFlagVersion,
		"version",
		"v",
		"",
		"New version",
	)
	Cmd.MarkFlagRequired("version")
	Cmd.Flags().BoolVarP(
		&CmdFlagNoCommit,
		"no-commit",
		"n",
		false,
		"Don't create commit with version bump",
	)
	Cmd.Flags().BoolVarP(
		&CmdFlagTag,
		"tag",
		"t",
		false,
		"Create also git tag",
	)
}
