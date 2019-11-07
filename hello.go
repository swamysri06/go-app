package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello Go-App</h1>")
}

func option1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Option1 check</h1>")
}

func option2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Option2 check</h1>")
}

func option3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Option3 check</h1>")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/option1", option1)
	http.HandleFunc("/option2", option1)
	http.HandleFunc("/option3", option1)
	fmt.Println("Server starting...")
	http.ListenAndServe(":3000", nil)
}

