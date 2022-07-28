package main

import "fmt"

func square(length, width float32) float32 {
	var Area = length * width
	return Area
}
func main() {
	s := square(5.3, 2.4)
	fmt.Println(s)

}
