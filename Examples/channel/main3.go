package main

import (
	"fmt"
	"time"
)

func main() {
	outPut1 := make(chan string)
	outPut2 := make(chan string)

	go server1(outPut1)
	go server2(outPut2)
	time.Sleep(time.Millisecond * 1)

	//outer:
	//for {
	select {
	case s1 := <-outPut1:
		fmt.Println(s1)
		//break outer
	case s2 := <-outPut2:
		fmt.Println(s2)
		//break outer
	default:
		fmt.Println("nothing happend")
	}
	//}

}

func server1(ch chan<- string) {
	fmt.Println("iam from the server1")
	ch <- "server1"
}
func server2(ch1 chan<- string) {
	fmt.Println("iam from the server2")
	ch1 <- "server2"
}
