package main

import (
	"fmt"
	"go-app/greetings"
)

func main() {
	// get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
