package slug_utils

import (
	"regexp"
	"strings"
)

var re = regexp.MustCompile("[^a-z0-9]+")

func Slugify(s string) string {
	return strings.Trim(re.ReplaceAllString(strings.ToLower(s), "-"), "-")
}

func SlugifyUnderscore(s string) string {
	return strings.Replace(Slugify(s), "-", "_", -1)
}
