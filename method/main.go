package main

import (
	"fmt"
	"method/name"
)

type Person struct {
	name.Animal
}

func (a Person) test() {
	fmt.Println(a.Name)
}

func main()  {
	var a Person
	a.Name = "tom"
	a.test()
}