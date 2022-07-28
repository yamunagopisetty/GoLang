package main

import "fmt"

//pointer is hold the address of variables

const (
	a = 2
	b
	c
)

func main() {
	var x = 5
	//var p *int
	p := &x
	//fmt.Printf("x = %d", p)
	defer fmt.Printf("x = %d", p)
	fmt.Println(a, b, c)
}
