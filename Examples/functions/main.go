package main

import "fmt"

func main() {
	school()

	result := details("yamuna", "Basinepalli")

	fmt.Println(result)

	prores := Prodetails(8, 4, 53, 5, 4, 6, 34, 4)
	fmt.Println("pro result is:", prores)

}

func school() {
	fmt.Println("welcome to our school")
}
func details(Name string, Address string) string {
	return Name + Address
}
func Prodetails(values ...int) int {
	total := 0

	for v := range values {
		total += v
	}

	return total
}
