package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "yamuna"
	fmt.Println(s)
	res := strings.Count(s, "a")
	res1 := strings.Count(s, "y")
	res2 := strings.Count(s, "m")
	fmt.Println("a is", res)
	fmt.Println("y is", res1)
	fmt.Println("m is", res2)
}
