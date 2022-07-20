package main


import(
	"fmt"
)



func main(){

	ch:=make(chan int)
	var (
		a,b=10,20
		
		//fmt.Println(c)
	)
	

	go add(ch)

	ch <- c
	//go mui()

}

func add(){

	
	
	
}

