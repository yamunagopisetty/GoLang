package main

import "fmt"

type Employee struct {
	Name   string
	Salary int
}

func (e Employee) EmployeeDetails() {
	fmt.Println("Name:", e.Name, "\nSalary:", e.Salary)
}

func ToString(e Employee) string {
	return fmt.Sprintln("Name:", e.Name, "\nSalary:", e.Salary)
}

func main() {
	Emp1 := Employee{"yamuna", 2345}
	Emp1.EmployeeDetails()

	Emp2 := Employee{"sai", 23455}
	Emp2.EmployeeDetails()

	Emp3 := Employee{"teja", 23345}
	//Emp2.EmployeeDetails()

	str := ToString(Emp3)
	fmt.Println(str)
}
