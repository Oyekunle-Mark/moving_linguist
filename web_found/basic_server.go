package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Path: ", r.URL.Path)
	fmt.Println("Scheme: ", r.URL.Scheme)

	fmt.Fprintf(w, "Hello, Chief!")
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":80", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
