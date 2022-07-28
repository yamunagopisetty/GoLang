package main

import (
	"fmt"
)

func times(arr []int, ch chan int) {
	for _, eleme := range arr {
		ch <- eleme * 3
	}
}
func timesthree(arr []int, ch chan int) {
	for _, elem := range arr {
		ch <- elem - 3
	}
}

func main() {

	arr := []int{1, 3, 4, 6, 8}
	ch := make(chan int, len(arr))
	ch1 := make(chan int, len(arr))

	go times(arr, ch)
	go timesthree(arr, ch1)
	//time.Sleep(time.Millisecond)

	for i := 0; i < len(arr)*2; i++ {
		select {
		case msg1 := <-ch:
			fmt.Println(msg1)
		case msg2 := <-ch1:
			fmt.Println(msg2)
		default:
			fmt.Println("Print channels")
		}
	}

}
