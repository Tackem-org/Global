package helpers

import (
	"math/rand"
	"time"
)

func RandStr(length int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.Seed(time.Now().UnixNano())
	r := make([]rune, length)
	for i := range r {
		r[i] = letters[rand.Intn(len(letters))]
	}
	key := string(r)
	return key
}
