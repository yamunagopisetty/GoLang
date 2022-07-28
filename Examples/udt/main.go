package main

import (
	"fmt"
	"udt/shape/rectangle"
	"udt/shape/square"
)

func main() {

	//shape.information()
	//fmt.Println()

	rct := rectangle.Rect{Length: 1.32, Bredth: 6.246}
	fmt.Println("the area of rectangle is", rct)

	sq := new(square.Square)

	sq.Side = 5.455

	fmt.Println(sq)
}
