package models

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) checkout() {
	fmt.Printf("My name is %v, I'm %v", p.Name, p.Age)
}
