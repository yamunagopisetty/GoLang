package main

import (
	"fmt"
)

func main() {
	var student []string

	student = append(student, "yamuna", "Gopisetty")

	//time.Sleep(5 * time.Second)

	fmt.Println(student)

	var (
		a, b, c = 1, 2, 3
	)

	a, b, c = b, c, a

	fmt.Println(a, b, c)

	//fmt.Println(student[a, b, c])
	fmt.Println(swap())

}

func swap() []int {
	i, j := 4, 5
	i, j = j, i
	return []int{i, j}
}
