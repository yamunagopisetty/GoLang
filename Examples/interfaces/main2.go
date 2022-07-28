package main

import "fmt"

type circle interface {
	Area() float32
	//Volume()float32
}
type values struct {
	radius float32
	//height float32
}

func (v *values) Area() float32 {
	return 3.14 * v.radius * v.radius
}
func main() {
	var c2 circle
	c := values{4.5}
	c2 = &c
	c2.Area()
	fmt.Println("Area of the circle is:", c2.Area())

}
