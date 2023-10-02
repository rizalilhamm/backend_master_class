package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Generate a Random Integer between Min and Max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(min+ max + 1)
}

// Generate a Random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// Generate a Random owner name
func RandomOwner() string {
	return RandomString(6)
}

// Generate random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{"IDR", "USD", "CAD"}
	n := len(currencies)
	
	return currencies[rand.Intn(n)]
}