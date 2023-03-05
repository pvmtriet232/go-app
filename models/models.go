package models

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) Checkout() {
	fmt.Printf("My name is %v, I'm %v \n", p.Name, p.Age)
}
