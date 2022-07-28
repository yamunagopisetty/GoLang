package main

import "fmt"

func main() {
	arr := "bhagya"
	fmt.Println("given string is:", arr)
	arr1 := []rune(arr)
	fmt.Println("duplicate elements in given string")
	for i := 0; i < len(arr1); i++ {
		for j := i + 1; j < len(arr1); j++ {
			if arr1[i] == arr1[j] {
				arr2 := string(arr1[j])
				fmt.Println(arr2)

			}

		}

	}

}
