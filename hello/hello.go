package main

import (
	"fmt"

	"github.com/pvmtriet232/go-app/greetings"
	_ "github.com/pvmtriet232/go-app/models"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
	Triet := &Person{"Triet", 26}
	Triet.checkout()

}
