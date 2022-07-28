package main

import "fmt"

type User struct {
	Name   string
	Number string
	Email  string
}

func main() {
	yamuna := User{"yamuna", "65398926", "yamuna@gmail.com"}
	fmt.Printf("the user struct details is :%+v\n", yamuna)

	user1 := &User{}
	user1.Name = "teja"
	user1.Number = "96367698"
	user1.Email = "teja@gmail.com"
	fmt.Printf("%+v", user1)

}
