package staged

import (
	"os"

	"github.com/go-git/go-git/v5"

	if_cmd "github.com/sikalabs/slu/cmd/git/if"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "staged",
	Short: "IF some files are staged",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		// os.Exit(1)
		r, err := git.PlainOpen(".")
		if err != nil {
			panic(err)
		}
		w, err := r.Worktree()
		if err != nil {
			panic(err)
		}
		s, err := w.Status()
		if err != nil {
			panic(err)
		}
		for _, f := range s {
			if f.Staging == git.Added {
				os.Exit(0)
			}
			if f.Staging == git.Modified {
				os.Exit(0)
			}
		}
		os.Exit(1)
	},
}

func init() {
	if_cmd.Cmd.AddCommand(Cmd)
}
