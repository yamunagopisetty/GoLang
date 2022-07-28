package main

import (
	"fmt"
	"time"
)

func main() {
	test := make(chan string)
	go school(test)
	go collage(test)
	time.Sleep(time.Millisecond * 1)
}
func school(name chan<- string) {
	fmt.Println("the name of the school is Zpgh school")
	name <- "mpp school"
}
func collage(name1 <-chan string) {
	test := <-name1
	fmt.Println(test)
}
