package helpers

import (
	"math/rand"
)

func RandString(length int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, length)
	lel := len(letters)
	for i := range b {
		b[i] = letters[rand.Intn(lel)]
	}
	return string(b)
}
