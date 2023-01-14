package git_utils

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

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
			if strings.HasPrefix(gitUrl, "ssh") {
				return strings.Replace(gitUrl, "ssh://git@", "https://", 1)
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

func TagNextCalver() {
	now := time.Now()
	nowY := now.Year()
	nowM := int(now.Month())
	r, err := git.PlainOpen(".")
	if err != nil {
		log.Fatalln(err)
	}
	tagsIter, err := r.Tags()
	if err != nil {
		panic(err)
	}
	tags := []string{}
	err = tagsIter.ForEach(func(b *plumbing.Reference) error {
		tags = append(tags, b.Name().Short())
		return nil
	})
	if err != nil {
		panic(err)
	}

	nextMicro := 0

	for _, tag := range tags {
		y, m, micro, err := ParseCalverTag(tag)
		if err != nil {
			continue
		}
		if y != nowY || m != nowM {
			continue
		}
		if micro >= nextMicro {
			nextMicro = micro + 1
		}
	}

	nextTag := fmt.Sprintf("v%d.%d.%d", nowY, nowM, nextMicro)
	err = exec_utils.ExecOut("git", "tag", nextTag)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(nextTag)
}

func ParseCalverTag(s string) (int, int, int, error) {
	r := regexp.MustCompile(`^v(\d{4}).(\d{1,2}).(\d+)$`)
	match := r.FindStringSubmatch(s)
	ok := len(match) == 4
	if !ok {
		return 0, 0, 0, fmt.Errorf("tag %s is not valid calver tag", s)
	}
	y, err := strconv.Atoi(match[1])
	if err != nil {
		return 0, 0, 0, err
	}
	m, err := strconv.Atoi(match[2])
	if err != nil {
		return 0, 0, 0, err
	}
	micro, err := strconv.Atoi(match[3])
	if err != nil {
		return 0, 0, 0, err
	}
	return y, m, micro, nil
}
