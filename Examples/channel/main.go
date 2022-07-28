package main

import (
	"fmt"
	"time"
)

func greet(done chan bool) {
	fmt.Println("good morning")
	time.Sleep(time.Millisecond * 1)
	done <- true
}

func main() {
	done := make(chan bool)
	go greet(done)
	<-done
	fmt.Println("hey")
}
