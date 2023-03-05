package main

import (
	"fmt"

	"github.com/pvmtriet232/go-app/greetings"
)

func main() {
	// get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
	Ann := &People{"Ann", 26}
	fmt.Println(Ann)
}
