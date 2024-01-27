package set_image

import (
	"os"
	"regexp"

	"github.com/go-git/go-git/v5"
	parent_cmd "github.com/sikalabs/slu/cmd/helm"

	"github.com/spf13/cobra"
)

var FlagVersion string
var FlagFile string
var CmdFlagNoCommit bool
var CmdFlagTag bool
var FlagScope string

func replace(s, key, value string) string {
	r := regexp.MustCompile(key + `: +([\w\./:_-]+)`)
	return r.ReplaceAllString(s, key+": "+value)
}

var Cmd = &cobra.Command{
	Use:     "version-bump",
	Short:   "Set version in Helm Chart",
	Aliases: []string{"vb"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		data, _ := os.ReadFile(FlagFile)
		s := replace(string(data), "version", FlagVersion)
		os.WriteFile(FlagFile, []byte(s), 0644)

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
		_, err = w.Add(FlagFile)
		if err != nil {
			panic(err)
		}

		scope := ""
		if FlagScope != "" {
			scope = "(" + FlagScope + ")"
		}

		commit, _ := w.Commit("VERSION"+scope+": "+FlagVersion, &git.CommitOptions{})
		_, err = r.CommitObject(commit)
		if err != nil {
			panic(err)
		}

		if CmdFlagTag {
			_, err := r.CreateTag(
				FlagVersion,
				commit,
				&git.CreateTagOptions{
					Message: "VERSION" + scope + ": " + FlagVersion,
				},
			)
			if err != nil {
				panic(err)
			}
		}
	},
}

func init() {
	parent_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagVersion,
		"version",
		"v",
		"",
		"Nev version",
	)
	Cmd.MarkFlagRequired("version")
	Cmd.Flags().StringVarP(
		&FlagFile,
		"file",
		"f",
		"",
		"Chart.yaml file",
	)
	Cmd.MarkFlagRequired("file")
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
	Cmd.Flags().StringVarP(
		&FlagScope,
		"scope",
		"s",
		"",
		"Add optional scope (for Conventional Commits)",
	)
}
