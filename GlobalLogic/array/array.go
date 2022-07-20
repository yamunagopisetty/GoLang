package main

import (
	"fmt"
)

func main() {
	var arr [3]int
	fmt.Println(arr)

	var arr1 [3]int
	arr1[0] = 100
	arr1[1] = 200
	arr1[2] = 300
	//arr1[3] = 400 since the length of the array is 3, you cannot assign 4th element
	fmt.Println(arr1)
	arr2 := [3]int{10, 11, 12}
	fmt.Println(arr2)

	for i := 0; i < len(arr2); i++ {

		fmt.Print(" ", arr2[i])
	}
}
