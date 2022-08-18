package semver_utils

import (
	"log"
	"strings"

	"github.com/aquasecurity/go-version/pkg/semver"
)

func CheckMinimumVersion(current, minimum string) bool {
	v1, err := semver.Parse(strings.Trim(current, "v"))
	if err != nil {
		log.Fatal(err)
	}

	v2, err := semver.Parse(strings.Trim(minimum, "v"))
	if err != nil {
		log.Fatal(err)
	}

	return v1.GreaterThanOrEqual(v2)
}
