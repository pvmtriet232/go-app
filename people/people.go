package people

import "fmt"

type People struct {
	Name string
}

func (p *People) Describe() {
	fmt.Printf("My name is %v.", p.Name)
}
