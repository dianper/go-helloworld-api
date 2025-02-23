package hello

import (
	"fmt" 
	"rsc.io/quote/v4"
)

func Hello(name string) string {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Hello())

	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}