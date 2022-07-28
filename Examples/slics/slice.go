package main

import "fmt"

func contains(s []string, str string) bool {
	for _, a := range s {
		if a == str {
			return true
		}
	}
	return false
}
func contains1(p []int, str int) bool {
	for _, b := range p {
		if b == str {
			return true
		}
	}
	return false
}

func main() {

	s := []string{"yamuna", "sai", "teja"}
	fmt.Println(contains(s, "yauna"))
	fmt.Println(contains(s, "sai"))
	fmt.Println(s, "yamuna")
	fmt.Println(s, "teja")

	p := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(contains1(p, 1))
	fmt.Println(contains1(p, 6))

	if p == nil {
		fmt.Println("number is present in slice")
	} else {
		fmt.Println("number is not present in slice")
	}
	//fmt.Println(p, 4)

	/*slice := []int{1, 2, 3, 4, 5}

	element := 1
	result := false
	for _, v := range slice {
		if v == element {
			result = true
			break
		}
	}
	if result {
		fmt.Print("element is present")
	} else {
		fmt.Print("element is present")
	}*/

}
