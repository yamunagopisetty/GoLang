package main

/*func foo() (string, string) {
	return "Yamuna", "Gopisetty"
}

func main() {
	fmt.Println(foo())

	s := "yamuna"

	//for _, v := range s {
	//defer fmt.Print(v)
	for i := 0; i < len(s); i++ {
		defer fmt.Print(s[i])

	}

}*/
/*func reverse(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {

		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)

}
func main() {
	fmt.Printf("%v\n", reverse("yamuna"))

}*/
func main() {
	str := "yamuna"
	chars := []rune(str)
	for _, n := range chars {
		yamu := string(n)
		defer print(yamu)

	}

}
