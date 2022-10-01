package utils

import (
	"math/rand"
	"time"
)

const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandomCode(n int) string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rnd.Int63()%int64(len(letters))]
	}
	return string(b)
}
