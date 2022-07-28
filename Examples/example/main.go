package main

import (
	"fmt"
	"strings"
)

func main() {
	var str strings.Builder

	for i := 0; i < 10; i++ {
		str.WriteString("a\n")
	}

	fmt.Println(str.String())
}
