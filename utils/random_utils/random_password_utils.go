package random_utils

import (
	"fmt"
	"math/rand"
	"strings"
	"unicode"
)

func RandomPassword() (string, error) {
	i := 0
	for {
		s := RandomString(16, ALL)
		if containsLowercase(s) && containsUpercase(s) && containsDigit(s) {
			return addUnderscores(s), nil
		}
		if i > 20 {
			return "", fmt.Errorf("cannot generate password")
		}
	}
}

func RandomHumanReadablePassword() string {
	words := make([]string, 4)
	for i := range words {
		words[i] = capitalize(BIP39_WORDS[rand.Intn(len(BIP39_WORDS))]) + fmt.Sprintf("%02d", rand.Intn(100))
	}

	result := strings.Join(words, "_")
	return result
}

func containsLowercase(s string) bool {
	for _, r := range s {
		if unicode.IsLower(r) {
			return true
		}
	}
	return false
}

func containsUpercase(s string) bool {
	for _, r := range s {
		if unicode.IsLower(r) {
			return true
		}
	}
	return false
}

func containsDigit(s string) bool {
	for _, r := range s {
		if unicode.IsDigit(r) {
			return true
		}
	}
	return false
}

func addUnderscores(input string) string {
	var blocks []string
	for i := 0; i < len(input); i += 4 {
		end := i + 4
		if end > len(input) {
			end = len(input)
		}
		blocks = append(blocks, input[i:end])
	}

	return strings.Join(blocks, "_")
}

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}
