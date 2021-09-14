package random_utils

import (
	"math/rand"
	"time"
)

const (
	LOWER  = "abcdefghijklmnopqrstuvwxyz"
	UPPER  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	DIGITS = "0123456789"
)

func RandomString(length int, chars string) string {
	if chars == "" {
		chars = LOWER + DIGITS
	}
	rand.Seed(time.Now().UnixNano())
	ll := len(chars)
	b := make([]byte, length)
	rand.Read(b) // generates len(b) random bytes
	for i := 0; i < length; i++ {
		b[i] = chars[int(b[i])%ll]
	}
	return string(b)
}
