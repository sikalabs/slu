package random_utils

import (
	"math/rand"
	"time"
)

const (
	LOWER  = "abcdefghijklmnopqrstuvwxyz"
	UPPER  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	DIGITS = "0123456789"
	ALL    = LOWER + UPPER + DIGITS
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

func RandomInt(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max+min) - min
}

func RandomBool() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2) == 0
}
