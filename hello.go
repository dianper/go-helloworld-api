package hello

import (
	"fmt" 
	"rsc.io/quote/v4"
)

func main(name string) string {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Hello())

	message := fmt.Sprintf("Hi, %v. Welcome!%", name)

	return 
}