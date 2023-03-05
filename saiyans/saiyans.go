package saiyans

import "fmt"

type Saiyans struct {
	Name  string
	Power int
	SSJ   bool
}

func (s *Saiyans) Checkout() {
	fmt.Printf("Name : %v, Power: %v , super saiyan status: %v", s.Name, s.Power, s.SSJ)
}
