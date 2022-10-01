package utils

import (
	"math/rand"
)

const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandomCode(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Int63()%int64(len(letters))]
	}
	return string(b)
}
