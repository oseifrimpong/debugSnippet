package main

import (
	"io"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloHandler(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Hello, World!")
}
