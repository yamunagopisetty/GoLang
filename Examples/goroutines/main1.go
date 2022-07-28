package main

import (
	"fmt"
	"time"
)

func display(message string) {
	fmt.Println(message)
}

func main() {
	go display("hello world")
	go display("my name is yamuna")
	go display("hey")
	go display("hello")
	time.Sleep(time.Millisecond * 1)

}
