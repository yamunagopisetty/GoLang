package main

import (
	"fmt"
	"time"
)

func incr(c chan int, x int) {
	c <- x + 1

	->ch
}

func main() {
	ch := make(chan int)
	//go func() {
	for i := 1; i <= 10; i++ {
		go incr(ch, i)
		ch <- i
		time.Sleep(time.Second * 1)

		//fmt.Println("The final value of X",1000)
	}

	fmt.Println("The final value of X", 1000)
    
	/*for item := range ch {
		fmt.Println(item)

	}*/

}
