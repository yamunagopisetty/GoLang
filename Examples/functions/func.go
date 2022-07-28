package main

import (
	"fmt"
	"math/rand"
)

func main() {

	fmt.Println("Random string")
	fmt.Println("11 chars :" + RandomString(11))
	fmt.Println("15 chars :" + RandomString(15))
}

func RandomString(n int) string {
	var chars = []rune("abcdefghijklmnopqrstuvwxyz")

	v := make([]rune, n)

	for i := range v {
		v[i] = chars[rand.Intn(len(chars))]
	}
	return string(v)
}
