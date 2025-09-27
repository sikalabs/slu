package upgrade_path

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strings"
)

// Detects installed edition (ce/ee) and returns semantic version "X.Y.Z".
func DetectGitLabEditionAndVersion() (edition, semver string, err error) {
	// Try CE then EE
	if v, e := getVersionFromPkg("gitlab-ce"); e == nil {
		return "ce", extractSemver(v), nil
	}
	if v, e := getVersionFromPkg("gitlab-ee"); e == nil {
		return "ee", extractSemver(v), nil
	}
	return "", "", errors.New("neither gitlab-ce nor gitlab-ee appears installed")
}

func getVersionFromPkg(pkg string) (string, error) {
	// Prefer apt (since user asked for apt list --installed)
	if v, err := versionFromAptList(pkg); err == nil && v != "" {
		return v, nil
	}
	// Fallback to dpkg-query for robustness
	if v, err := versionFromDpkg(pkg); err == nil && v != "" {
		return v, nil
	}
	return "", fmt.Errorf("could not determine version for %s", pkg)
}

func versionFromAptList(pkg string) (string, error) {
	cmd := exec.Command("apt", "list", "--installed", pkg)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out // capture "Listing..." etc
	if err := cmd.Run(); err != nil {
		return "", err
	}

	// Examples of possible lines:
	// gitlab-ce/jammy,now 18.0.1-ce.0 amd64 [installed]
	// gitlab-ee/unknown,now 17.9.2-ee.0 amd64 [installed]
	// Some apt versions print "Listing... Done" on the first line.
	lines := strings.Split(out.String(), "\n")
	re := regexp.MustCompile(pkg + `/.*now ([^ ]+)`) // capture the version token
	for _, ln := range lines {
		if m := re.FindStringSubmatch(ln); len(m) == 2 {
			return m[1], nil
		}
	}
	return "", errors.New("no apt list match")
}

func versionFromDpkg(pkg string) (string, error) {
	// Prints raw version like: 18.0.1-ce.0
	cmd := exec.Command("dpkg-query", "-W", "-f=${Version}\n", pkg)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	if err := cmd.Run(); err != nil {
		return "", err
	}
	return strings.TrimSpace(out.String()), nil
}

// Extracts "X.Y.Z" prefix from things like "18.0.1-ce.0" or "18.0.1+omnibus-1"
func extractSemver(v string) string {
	re := regexp.MustCompile(`\b(\d+\.\d+\.\d+)\b`)
	m := re.FindStringSubmatch(v)
	if len(m) >= 2 {
		return m[1]
	}
	// Fallback: return original if pattern not found
	return v
}

func BuildUpgradeURL(edition, semver string) string {
	return fmt.Sprintf(
		"https://gitlab-com.gitlab.io/support/toolbox/upgrade-path/?current=%s&edition=%s",
		semver, edition,
	)
}

func GetGitlabUpgradePathURL() string {
	ed, ver, err := DetectGitLabEditionAndVersion()
	if err != nil {
		log.Fatalf("Error detecting Gitlab edition and version: %v", err)
	}
	url := BuildUpgradeURL(ed, ver)
	return url
}
