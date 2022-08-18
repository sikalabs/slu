package git_utils

import (
	"log"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/sikalabs/slu/utils/exec_utils"
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
			if strings.HasPrefix(gitUrl, "https") {
				return gitUrl
			}
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

func DeleteBranch(name string) {
	err := exec_utils.ExecOut("git", "branch", "-D", name)
	if err != nil {
		log.Fatalln(err)
	}
}

func DeleteAllDependabotBranches() {
	branches := GetLocalBranches()
	for _, branch := range branches {
		if strings.HasPrefix(branch, "dependabot/") {
			DeleteBranch(branch)
		}
	}
}

func GetCurrentBranch() string {
	r, err := git.PlainOpen(".")
	if err != nil {
		panic(err)
	}

	head, err := r.Head()
	if err != nil {
		panic(err)
	}

	return head.Name().Short()
}

func DeleteAllLocalBranches() {
	branches := GetLocalBranches()
	currentBranch := GetCurrentBranch()
	for _, branch := range branches {
		if branch == currentBranch {
			continue
		}
		DeleteBranch(branch)
	}
}

func UseSSH() {
	r, err := git.PlainOpen(".")
	if err != nil {
		log.Fatalln(err)
	}
	remotes, err := r.Remotes()
	if err != nil {
		log.Fatalln(err)
	}
	remoteName := remotes[0].Config().Name
	oldUrl := remotes[0].Config().URLs[0]
	newUrl := strings.Replace(oldUrl, "https://", "ssh://git@", 1)
	r.DeleteRemote(remoteName)
	r.CreateRemote(&config.RemoteConfig{
		Name: remoteName,
		URLs: []string{newUrl},
	})
}
