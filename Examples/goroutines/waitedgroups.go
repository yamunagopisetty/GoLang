package main

import (
	"fmt"
	"sync"
	"time"
)

func fruits(things string) {
	//for i := 0; i <= 5; i++ {
	fmt.Println(things)
	time.Sleep(time.Millisecond * 500)
	//}//

}

func main() {

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		fruits("apple")
		wg.Done()
	}()
	go func(a, b int) {
		for i := a; i < b; i++ {
			fruits("mango")
			time.Sleep(time.Millisecond * 1)
			wg.Done()
		}
	}(1, 3)
	go func() {
		fruits("banana")
		wg.Done()
	}()
	wg.Wait()
}
