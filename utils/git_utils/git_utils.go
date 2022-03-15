package git_utils

import (
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func GetRepoUrl() string {
	r, err := git.PlainOpen(".")
	if err != nil {
		panic(err)
	}
	remotes, err := r.Remotes()
	if len(remotes) > 0 && err == nil {
		remote := remotes[0]
		if len(remote.Config().URLs) > 0 {
			gitUrl := remote.Config().URLs[0]
			x := strings.Replace(gitUrl, ":", "/", 1)
			x = strings.Replace(x, ".git", "", 1)
			url := strings.Replace(x, "git@", "https://", 1)
			return url
		}
	}
	return ""
}

func GetNewAddedFiles(repoPath string) ([]string, error) {
	var newFiles []string
	r, err := git.PlainOpen(repoPath)
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
	for name, f := range s {
		if f.Staging == git.Added {
			newFiles = append(newFiles, name)
		}
	}
	return newFiles, nil
}

func GetLocalBranches() []string {
	var branches []string
	r, err := git.PlainOpen(".")
	if err != nil {
		panic(err)
	}
	branchesIter, err := r.Branches()
	if err != nil {
		panic(err)
	}
	err = branchesIter.ForEach(func(b *plumbing.Reference) error {
		branches = append(branches, b.Name().Short())
		return nil
	})
	if err != nil {
		panic(err)
	}
	return branches
}
