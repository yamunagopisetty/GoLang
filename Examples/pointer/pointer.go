package main

import "fmt"

func change(val *int) {
	*val = 34
}
func changestring(s *string) {
	*s = "yamuna"
}
func modify(arr *[3]int) {
	arr[0] = 5

}
func modifys(i []string) {
	if len(i) > 0 {
		i[0] = "ananya"
	}
}
func modify2(i []int) {
	if len(i) > 0 {
		i[3] = 7
	}
}

func main() {
	a := 23
	fmt.Println("before change:", a)
	b := &a
	change(b)
	fmt.Println("after change:", a)

	p := "sai"
	fmt.Println("before change:", p)
	m := &p
	changestring(m)
	fmt.Println("after change:", p)

	array := [3]int{2, 3, 4}
	fmt.Println("before change:", array)
	//fmt.Println(array)
	ar := &array
	modify(ar)
	fmt.Println("after change:", array)

	slic := []string{"yamuna", "sai", "teja", "bhargav", "divya"}
	fmt.Println("before change:", slic)
	//slic[1] = "pranavi"

	//i := &slic
	//modifys(i)
	fmt.Println("after change:", slic)

	slic1 := []int{3, 5, 78, 98, 67}
	fmt.Println("before change:", slic1)
	//slic[1] = "pranavi"

	modify2(slic1)
	fmt.Println("after change:", slic1)

}
