package main

import (
	"fmt"
	"go-helloworld-api/hello"
)

func main() {
	message := hello.Hello("Diego")
	fmt.Println(message)
}