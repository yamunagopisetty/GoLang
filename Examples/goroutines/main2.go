package main

import (
	"fmt"
	"time"
)

func main() {
	go greet()

	fn := func() {
		for i := 1; i <= 10; i++ {
			go fmt.Println("Hello Anonymous function")
		}
	}

	go fn()
	// anonymous function
	go func(si, ei int) {
		for i := si; i <= ei; i++ {
			go greetI(i)
		}
	}(1, 100)

	time.Sleep(time.Millisecond * 10)
	fmt.Println("Hello World")
}

func greet() {
	go fmt.Println("Hello Go routine")

	go func() {
		fmt.Println("----> Hello Go routine <---------")
	}()
}

func greetI(i int) {
	fmt.Println("Hello Go routine--->", i)
}
