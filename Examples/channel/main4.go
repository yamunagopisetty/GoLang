package main

import (
	"fmt"
	"time"
)

func main() {
	addition := make(chan int)
	go add(addition)
	//<-addition
	time.Sleep(time.Millisecond * 1)
}
func add(ch chan<- int) {
	fmt.Println("addition of two numbers is")
	a := 4 + 5
	//fmt.Println()
	ch <- a

}
