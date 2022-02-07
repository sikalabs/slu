package git_utils

import (
	"strings"

	"github.com/go-git/go-git/v5"
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
