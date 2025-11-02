package version_bump

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

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
var CmdFlagAuto bool

var Cmd = &cobra.Command{
	Use:     "version-bump",
	Short:   "Bumb & commit version of (SL) Go application",
	Aliases: []string{"vb"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		// Handle --auto flag
		if CmdFlagAuto {
			currentVersionBytes, err := os.ReadFile("version/version.go")
			if err != nil {
				log.Fatalln("Failed to read current version:", err)
			}

			currentVersionContent := string(currentVersionBytes)
			// Parse current version from: var Version string = "vX.Y.Z-dev"
			start := strings.Index(currentVersionContent, `"`)
			end := strings.LastIndex(currentVersionContent, `"`)
			if start == -1 || end == -1 || start == end {
				log.Fatalln("Failed to parse current version from version/version.go")
			}

			currentVersion := currentVersionContent[start+1 : end]
			CmdFlagVersion = calculateNextVersion(currentVersion)
			fmt.Printf("Auto-bumping version: %s -> %s\n", currentVersion, CmdFlagVersion)
		}

		// Validate that --tag is not used with -dev versions
		if CmdFlagTag && strings.HasSuffix(CmdFlagVersion, "-dev") {
			log.Fatalln("Cannot use --tag flag with -dev version:", CmdFlagVersion)
		}

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

func calculateNextVersion(currentVersion string) string {
	// If version ends with -dev, remove it (e.g., v0.1.0-dev -> v0.1.0)
	if strings.HasSuffix(currentVersion, "-dev") {
		return strings.TrimSuffix(currentVersion, "-dev")
	}

	// Otherwise, bump minor version and add -dev (e.g., v0.1.0 -> v0.2.0-dev)
	// Parse version: vX.Y.Z
	version := strings.TrimPrefix(currentVersion, "v")
	parts := strings.Split(version, ".")
	if len(parts) != 3 {
		log.Fatalln("Invalid version format:", currentVersion)
	}

	major := parts[0]
	minor, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatalln("Failed to parse minor version:", err)
	}

	return fmt.Sprintf("v%s.%d.0-dev", major, minor+1)
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
	Cmd.Flags().BoolVarP(
		&CmdFlagAuto,
		"auto",
		"a",
		false,
		"Auto-calculate next version (v0.1.0-dev -> v0.1.0, v0.1.0 -> v0.2.0-dev)",
	)
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
		&CmdFlagGoGit,
		"go-git",
		false,
		"Use go-git instead of git binary",
	)
	Cmd.MarkFlagsOneRequired("version", "auto")
	Cmd.MarkFlagsMutuallyExclusive("version", "auto")
}
