package main

import "fmt"

type User struct {
	Name   string
	Status bool
}

func (u User) getStatus() {
	fmt.Println("is user active:", u.Status)

}
func (u User) getName() {
	fmt.Println("is user Name is:", u.Name)

}

func main() {
	user1 := User{"yamuna", true}
	user1.getStatus()
	user1.getName()
}
