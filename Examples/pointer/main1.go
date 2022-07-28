package main

import "fmt"

func main() {
	var fun string = "yamuna"
	var pointer *string = &fun

	fmt.Println("func = ", fun)
	fmt.Println("pointer =", pointer)

	//fmt.Println("*pointer", *pointer)

	*pointer = "sai"
	fmt.Println("*pointer", *pointer)

	fmt.Println("func = ", fun)
	var name Name = Name{ts: "teja"}
	fmt.Println(name)
	changeName(&name)
}

type Name struct {
	ts string
}

func changeName(name *Name) {
	name.ts = "bharahav"
	fmt.Println(name)
}
