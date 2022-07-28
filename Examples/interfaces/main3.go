package main

import "fmt"

//type assortion in interfaces
func types(i interface{}) {
	v, _ := i.(int) //------>type assortion i.()
	fmt.Println(v)
}
func main() {
	s := 6
	types(s)

}
