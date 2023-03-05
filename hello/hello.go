package main

import (
	"fmt"

	"github.com/pvmtriet232/go-app/greetings"
	"github.com/pvmtriet232/go-app/models"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
	Triet := &models.Person{Name: "Triet", Age: 26}
	Triet.Checkout()

}
