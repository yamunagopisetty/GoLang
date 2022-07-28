package main

import (
	"fmt"
	"sync"
	"time"
)

func greet(wg *sync.WaitGroup) {
	fmt.Println("hello world")
	time.Sleep(time.Millisecond * 1)
	wg.Done()
}
func main() {

	wg := sync.WaitGroup{}

	for i := 1; i < 3; i++ {
		wg.Add(1)
		go greet(&wg)
		time.Sleep(time.Millisecond * 1)

	}
	fmt.Println("hey")
	wg.Wait()

}
