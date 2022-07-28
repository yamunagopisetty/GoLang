package main

import "fmt"

type rectangle interface {
	Area() float32
}
type rect struct {
	length, width float32
}

func (r rect) Area() float32 {
	return r.length * r.width
}
func main() {

	re := rect{3, 4}
	fmt.Println("the area of rectangle is", re.Area())

}
