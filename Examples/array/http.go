package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("server started...")
	http.HandleFunc("/greet", Greet)

	http.HandleFunc("/ping", Message("pong"))
	http.HandleFunc("/gopisetty", Messages("yamuna"))
	http.ListenAndServe(":9090", nil)
}

func Messages(messages string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, messages)
	}
}

func Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func Message(message string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, message)
	}
}
