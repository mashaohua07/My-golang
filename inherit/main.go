package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age int
	height float64
}

type tool struct {
	node bool
	str byte
}

type Simulation struct {
	Person
	Nide tool
	function string
}

func (t *Simulation) Nite() {
	fmt.Println(t.Name, t.Nide.node, t.function)
}

func main()  {

	rs := Simulation{
		Person: Person{Name: "ll",
		},

		Nide: tool{
			true,
			15,
		},
		function: "ool", }
	rs.Nite()
}
