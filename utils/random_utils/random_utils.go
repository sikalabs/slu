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
	ll := len(chars)
	b := make([]byte, length)
	rand.Read(b) // generates len(b) random bytes
	for i := 0; i < length; i++ {
		b[i] = chars[int(b[i])%ll]
	}
	return string(b)
}

func RandomInt(min int, max int) int {
	return randomIntn(max+min) - min
}

func RandomBool() bool {
	return randomIntn(2) == 0
}
