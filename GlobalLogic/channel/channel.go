package main

import (
	"fmt"
	//"time"
)

func main() {

	ch := make(chan int)

	go func() {
		for i := 1; i <= 100; i++ {
			ch <- i
		}
		close(ch)

	}()
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()

	/*output1 := make(chan string)
		output2 := make(chan string)

		go server1(output1)
		go server2(output2)

	outer:

		for {

			select {
			case s1 := <-output1:
				fmt.Println(s1)
				break outer

			case s2 := <-output2:
				fmt.Println(s2)
				break outer
			default:
				fmt.Println("Nothing happend")
			}
		}
	}

	func server1(ch chan string) {
		ms := int64(rand.Intn(10))
		time.Sleep(time.Millisecond * time.Duration(ms))
		ch <- "from server1"
	}

	func server2(ch chan string) {
		ms := int64(rand.Intn(10))
		time.Sleep(time.Millisecond * time.Duration(ms))
		ch <- "from server2"
	}*/

	/*var mychan chan int
	  fmt.Println("value of channel", mychan)

	  mychan1 := make(chan int)

	  fmt.Println("value of channel", mychan1)

	  data := make(chan string)
	  data <- "yamuna"
	  data <- "sai"
	  fmt.Println(<-data)

	  data <- "How are you doing"
	  fmt.Println(<-data)
	  fmt.Println(<-data)*/
	/*fmt.Println("start main method")

	  ch := make(chan int)

	  go myfunc(ch)

	  ch <- 40

	  fmt.Println("end main method")*/

	/*c := make(chan string)

	  go myfunc(c)

	  for {
	  	res, ok := <-c
	  	if ok == false {
	  		fmt.Println("channel close", ok)
	  		break
	  	}
	  	fmt.Println("channel open", res, ok)

	  }*/

	/*done := make(chan bool, 2)
	  go Employee(done)
	  <-done*/

	//func myfunc(ch chan int) {

	//	fmt.Println(50 - <-ch)//(50+<-ch)

	//}

	/*func myfunc(ch chan string) {

		for i := 1; i <= 5; i++ {
			ch <- "yamuna"
		}
		close(ch)
	}*/

	//func Employee(done chan bool) {
	//fmt.Println("Employee are working")
	//time.Sleep(time.Millisecond * 2)
	//fmt.Println("done")
	//done <- true*?
	//}
}
