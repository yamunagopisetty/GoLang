package main

import (
	"fmt"
	"sort"
)

func main() {
	slice := []int{1, 3, 5, 2, 4, 6}
	fmt.Println(slice)
	sort.Sort(sort.Reverse(sort.IntSlice(slice)))
	fmt.Println(slice)

	name := "yamuna"
	fmt.Println(name)
	//fmt.Printf("original string%s\n:", string(name))
	//fmt.Printf("reverse string:")
	for _, v := range []rune(name) {
		defer fmt.Printf("%c", v)
	}
}
