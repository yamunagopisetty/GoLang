package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Printf("%d.%d.%d\n", t.Day(), t.Month(), t.Year())
	fmt.Println(t.Date())
	//fmt.Println(t.Month(), t.Year())

}
