package main

import "fmt"

func main() {
	arr := [5]string{"yamuna", "sai", "teja"}
	arr1 := arr
	fmt.Println(arr)
	fmt.Println(arr1)
	arra := [5]int{1, 3, 5, 4, 6}
	arr2 := &arra
	fmt.Println(arra)
	fmt.Println(*arr2)
}
