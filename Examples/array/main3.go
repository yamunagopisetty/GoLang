package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "Race"
	s2 := "Care"
	s1 = strings.ToLower(s1)
	s2 = strings.ToLower(s2)

	char := []rune(s1)

	for i := 0; i < len(char); i++ {
		char := string(char[i])
		fmt.Print(char)
	}

	fmt.Println()
	char2 := []rune(s2)

	for i := 0; i < len(char2); i++ {
		char2 := string(char2[i])
		fmt.Print(char2)

	}

}
