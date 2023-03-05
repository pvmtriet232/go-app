package people

import "fmt"

type People struct {
	Name string
	Age  int
}

func (p *People) Describe() {
	fmt.Printf("My name is %v and I'm %v years old.", p.Name, p.Age)
}
