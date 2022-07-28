package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 2, 1}
	fmt.Println(arr)
	fmt.Println("duplicate elements in given array")
	for i := 0; i < 5; i++ {
		for j := i + 1; j < 5; j++ {
			if arr[i] == arr[j] {
				fmt.Println(arr[j])
			}
		}

	}

}
