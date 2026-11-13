package random_utils

import (
	"crypto/rand"
	"math/big"
)

const (
	LOWER  = "abcdefghijklmnopqrstuvwxyz"
	UPPER  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	DIGITS = "0123456789"
	ALL    = LOWER + UPPER + DIGITS
)

// randomIntn returns a cryptographically secure random int in [0, n).
func randomIntn(n int) int {
	v, err := rand.Int(rand.Reader, big.NewInt(int64(n)))
	if err != nil {
		panic(err)
	}
	return int(v.Int64())
}

func RandomString(length int, chars string) string {
	if chars == "" {
		chars = LOWER + DIGITS
	}
	b := make([]byte, length)
	for i := 0; i < length; i++ {
		b[i] = chars[randomIntn(len(chars))]
	}
	return string(b)
}

func RandomInt(min int, max int) int {
	return min + randomIntn(max-min)
}

func RandomBool() bool {
	return randomIntn(2) == 0
}
