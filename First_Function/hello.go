package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Nikhil")
	fmt.Println(message)
}
