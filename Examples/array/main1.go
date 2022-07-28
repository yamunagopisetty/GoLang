package main

import "fmt"

func main() {
	arr := [4]string{"yamuna", "sai", "yamuna", "divya"}
	fmt.Println(arr)
	fmt.Println("duplicate elements in given array")
	for i := 0; i < 4; i++ {
		for j := i + 1; j < 4; j++ {
			if arr[i] == arr[j] {
				fmt.Println(arr[j])
			}

		}
	}

}
