// Package shorten contains utility functions for shortening a URL.
package shorten

import (
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_-.")

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandSeq generate a random sequence based on the supplied length
func RandSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
