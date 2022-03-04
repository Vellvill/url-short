package utils

import (
	"math/rand"
	"time"
)

var crypto []rune = []rune("1234567890qwertyuiopQWERTYUIIOPasdfghjklASDFGHJKLzxcvbnmZXCVBNM")

func Encode(s []rune) string {
	r := []rune{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		r = append(r, crypto[rand.Intn(len(crypto)-0)+0])
	}

	return string(r)
}
