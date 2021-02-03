package random

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var stringRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

var integerRunes = []rune ("123456789")

func randomString (n int, runes []rune) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = runes[rand.Intn(len(runes))]
    }
    return string(b)
}

// String returns rand string
func String (length int) string {
	return randomString(length, stringRunes)
}

// Int returns rand string
func Int (length int) string {
	return randomString(length, integerRunes)
}