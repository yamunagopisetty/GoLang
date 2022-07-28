package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("started main thread")
	go long()
	go short()
	time.Sleep(2 * time.Millisecond)
	fmt.Println("ended mainthread")
}
func long() {
	fmt.Println("started long thread")
	time.Sleep(time.Millisecond)
	fmt.Println("ended long thread")

}
func short() {
	fmt.Println("started short thread")
	time.Sleep(time.Millisecond)
	fmt.Println("ended short thread")

}
