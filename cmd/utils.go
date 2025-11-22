package main

import (
	"math/rand"
	"time"
)

func random_string_from_charset(length int) string {

	charset := "abcdefghijklmnopqrstuvwxyz"
	
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	id := ""

	for i := 0; i < length; i++ {
		id += string(charset[r.Int() % len(charset)])
	}

	return id
}
