package main

import (
	"fmt"

	"github.com/pvmtriet232/go-app/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)

}
