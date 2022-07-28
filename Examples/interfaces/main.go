package main

import "fmt"

type washingmechine interface {
	cleaning()
	drying()
}

func cleaning() {
	fmt.Println("washing mechine is cleaning the cloths")
}
func drying() {
	fmt.Println("washing mechine is drying the cloths")
}

func main() {
	fmt.Println("------------washingmoechine activities are ------------")
	cleaning()
	drying()

}
