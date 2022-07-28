package main

import "fmt"

func main() {
	done := make(chan string, 4)
	done <- "yamuna"
	done <- "sai"
	done <- "teja"
	done <- "bhargav"

	fmt.Println(<-done, "\n", <-done, "\n", <-done, "\n", <-done)

}
