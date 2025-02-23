package main

import (
	"fmt"
	"go-helloworld-api/hello"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	message := hello.Hello("Diego")
	fmt.Fprintf(w, message)
}

func main() {
	http.HandleFunc("/", helloHandler)

	fmt.Println("Starting server on :8080..")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}