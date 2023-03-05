package main

import (
	"fmt"

	"github.com/pvmtriet232/go-app/greetings"
	_ "github.com/pvmtriet232/go-app/people"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
	Ann := People{"Ann", 25}
	fmt.Println(Ann)
}
