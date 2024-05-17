package version_bump

import (
	"log"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	go_code_cmd "github.com/sikalabs/slu/cmd/go_code"
	"github.com/sikalabs/slu/utils/exec_utils"

	"github.com/spf13/cobra"
)

var CmdFlagVersion string
var CmdFlagNoCommit bool
var CmdFlagTag bool
var CmdFlagGoGit bool

var Cmd = &cobra.Command{
	Use:     "version-bump",
	Short:   "Bumb & commit version of (SL) Go application",
	Aliases: []string{"vb"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		version_go_file := `package version

var Version string = "` + CmdFlagVersion + `"
`
		err := os.WriteFile("version/version.go", []byte(version_go_file), 0644)
		if err != nil {
			panic(err)
		}

		if CmdFlagNoCommit {
			return
		}

		var r *git.Repository
		var commit plumbing.Hash

		if !CmdFlagGoGit {
			err = exec_utils.ExecOut("git", "add", "version/version.go")
			if err != nil {
				log.Fatalln(err)
			}
			err = exec_utils.ExecOut("git", "commit", "-m", "VERSION: "+CmdFlagVersion)
			if err != nil {
				log.Fatalln(err)
			}
		} else {
			r, err = git.PlainOpen(".")
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
			commit, _ = w.Commit("VERSION: "+CmdFlagVersion, &git.CommitOptions{})
			_, err = r.CommitObject(commit)
			if err != nil {
				panic(err)
			}
		}

		if !CmdFlagGoGit {
			if CmdFlagTag {
				err = exec_utils.ExecOut("git", "tag", CmdFlagVersion, "-m", "VERSION: "+CmdFlagVersion)
				if err != nil {
					log.Fatalln(err)
				}
			}
		} else {
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
	Cmd.Flags().BoolVar(
		&CmdFlagNoCommit,
		"go-git",
		false,
		"Use go-git instead of git binary",
	)
}
