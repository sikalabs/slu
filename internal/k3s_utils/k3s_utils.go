package k3s_utils

import (
	"os"
	"path/filepath"
)

func CheckIfKubectlIsLinkOfK3s() bool {
	kubectl := "/usr/local/bin/kubectl"
	k3s := "/usr/local/bin/k3s"

	info, err := os.Lstat(kubectl)
	if err != nil {
		return false
	}

	if info.Mode()&os.ModeSymlink != 0 {
		_, err := os.Readlink(kubectl)
		if err != nil {
			return false
		}

		// Resolve both to absolute paths
		resolvedkubectl, err := filepath.EvalSymlinks(kubectl)
		if err != nil {
			return false
		}
		resolvedk3s, err := filepath.EvalSymlinks(k3s)
		if err != nil {
			return false
		}

		if resolvedkubectl == resolvedk3s {
			return true
		}
	}
	return false
}
