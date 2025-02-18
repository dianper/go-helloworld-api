package main

import (
	"fmt"
	"example/hello"
)

func main() {
	message := hello.Hello("Diego")
	fmt.Println(message)
}