package generator

import (
	crypto "crypto/rand"
	"encoding/hex"
	"math/rand"
	"sync"
	"time"
)

var mutex sync.Mutex
var src = rand.New(rand.NewSource(time.Now().UnixNano()))

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func int63() int64 {
	mutex.Lock()
	v := src.Int63()
	mutex.Unlock()
	return v
}

// RandString constructs a pseudo random string of length n
func RandString(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

// GenerateSecureString creates a cryptographically secure string of size n.
func GenerateSecureString(n int) (string, error) {
	b, err := generateRandomBytes(n)
	return hex.EncodeToString(b), err
}

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := crypto.Read(b)

	if err != nil {
		return nil, err
	}
	return b, nil
}

// RandomStrings generates n random strings of length 300
func RandomStrings(n int) []string {
	words := make([]string, n)
	for i := 0; i < n; i++ {
		words[i] = RandString(300)
	}
	return words
}

// RandomSimilarStrings returns strings with some probabilistic overlap.
func RandomSimilarStrings(n int, length int) []string {
	words := make([]string, n)
	for i := 0; i < n; i++ {
		words[i] = RandString(length)
	}
	return words
}
