package util

import (
	"math/rand"
	"time"
)

//generate user's name if he doesn't have one
func RandomName(n int) string {
	var letters = []byte("asdfghjklzxcvbnmQWERTYUIOPZXCVBNM")
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
