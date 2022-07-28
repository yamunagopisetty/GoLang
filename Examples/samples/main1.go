package main

import (
	"errors"
	"fmt"
)

func main() {
	answer, err := divided(2, 0)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(answer)
	}
}

func divided(x int, y int) (int, error) {
	if y == 0 {

		return -1, errors.New("the error are occured")
	}
	return x / y, nil
}
