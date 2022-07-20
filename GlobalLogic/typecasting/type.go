package main

import "fmt"

func main() {

	var long int64
	var small int32 = 32345

	long = int64(small) // type casting in go

	fmt.Println(long, small)
	fmt.Println("hey welcome")

	greeting()
	fmt.Println("Hello")
}

func greeting() {
	fmt.Println("good morning")
}
