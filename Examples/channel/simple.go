package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	go greet(done)

	data := make(chan string)
	go publish(data)
	go subscriber(data)

	time.Sleep(time.Millisecond * 1)

}
func greet(done chan<- bool) {
	fmt.Println("hello world")
	done <- true
}
func publish(send chan<- string) {
	fmt.Println("iam bublisher ....to publish the data")
	send <- "hello subscriber,from publisher"
}
func subscriber(reciver <-chan string) {
	data := <-reciver
	fmt.Println(data, "iam a subscriber recived the data")
}
