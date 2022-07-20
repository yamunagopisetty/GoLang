package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("D:/Yamuna.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		//fmt.Println(f.Name())
		yamu := make([]byte, 3)
		for {
			_, err := f.Read(yamu)
			if err != nil {
				break
			}
			fmt.Print(string(yamu))
		}

	}

}
