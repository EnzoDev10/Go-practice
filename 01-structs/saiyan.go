package main

import "fmt"

// Structs are similar to objects
type Person struct {
	Name string
}

// Structs can have functions that are related to parts of them. They are called receivers.
func (p *Person) Introduce() {
	fmt.Printf("Hi, I'm %s\n", p.Name)
}

// Structs can have other structs inside
type Saiyan struct {
	*Person
	Power int
}

func (s *Saiyan) Super() {
	s.Power += 10000
}
